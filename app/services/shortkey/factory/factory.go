package factory

import (
	"context"
	"database/sql"
	"eshort/app/biz"
	"eshort/app/models/eshort"
	"eshort/pkg/array_tool"
	"eshort/pkg/config"
	"eshort/pkg/ehash"
	"eshort/pkg/eredis"
	"eshort/pkg/model"
	"eshort/pkg/mtime"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"math/rand"
	"strconv"
)

var rate float64
var retry int
var SkeyPollMax uint64

func Setup() {
	rate = cast.ToFloat64(config.Get("eshort.grow"))
	retry = config.GetInt("eshort.clash_retry")
	SkeyPollMax = uint64(config.GetUint("eshort.pool_max"))
}

var SkeyContainerName = "eshort:pool:name"
var produceStatus = "eshort:produce:status"

// 单条不入缓存
func MakeRowSkey() (string, error) {
	maxId := getMaxId()
	skey := ehash.HashKey(maxId + 1)
	shortUrl := eshort.Eshort{ShortKey: skey, Ext: biz.BIZ.GetExt()}
	err := model.DB.Create(&shortUrl).Error
	if err != nil {
		return "", err
	}
	//go BatchProductionSkey(int64(SkeyPollMax))
	return skey, nil
}

// BatchProductionSkey 批量生产可用短链接key
func BatchProductionSkey(count int64) {
	//再次计算补充多少到skeypool
	//根据状态判断当前是否处于生产期，1为生产期。减少加锁解锁操作。
	//先获取状态，写时加锁，
	m := mtime.Mtime{}
	if inProduction() {
		fmt.Println("BatchProductionSkey", "状态为生产中，退出生产")
		return
	}
	key := "eshort:produce:lock"
	uskey := rand.Int63()
	eredis.GetLock(key, uskey, 4) //可能导致在并发中，第一批到达此的线程在此自旋。
	// 当线程A生产完后，释放锁，第一批的线程会继续在此生产。
	//有并发到达这一步，则说明有可能说明扩容需求大。但生产过程可能耗时，所以
	//需要有一个值控制最大长度，防止并发攻击，导致key池无限增长。默认长度，扩容长度，最大长度。
	//

	//如果选择的是主动触发扩容，就不需要开启监控线程。
	//如果选择的是被动触发扩容，就不会有主动触发线程。故这里不再加锁。
	bg := context.Background()
	eredis.RedisClient.Set(bg, produceStatus, 1, 0)
	growpool(SkeyPollMax - uint64(count))
	eredis.RedisClient.Set(bg, produceStatus, 0, 0)
	eredis.Unlock(key, uskey)
	fmt.Println("生产结束", m.GetFormatTime())
	return
}

// 是否需要扩容
func TooLittle(count int64) bool {
	need := int64(float64(SkeyPollMax) * rate)
	if count < need { //低于阈值，就去生产
		return true
	}
	return false
}

// growpool 生成多条可用短链接 key 入库并进缓存
func growpool(num uint64) {
	maxId := getMaxId()
	i := 1
	for {
		shortUrls, keys := hashKeys(maxId, num)
		err := model.DB.Create(&shortUrls).Error
		si := strconv.Itoa(i)
		fmt.Println("尝试批量插入" + si + "次")
		if err == nil {
			fmt.Println("在第" + si + "次插入时，插入成功")
			re := eredis.RedisClient.RPush(context.Background(), SkeyContainerName, keys...).Err()
			if re != nil {
				fmt.Println("数据库插入成功，缓存未更新成功")
			}
			break
		} else {
			i += 1
			fmt.Println("批量插入数据库，可能存在冲突，将再次生成key，尝试批量插入")
			fmt.Println("retry:", retry)
			if i >= retry {
				fmt.Println("可能发生严重错误需要排查")
				break
			}
		}
	}
	fmt.Println("生成结束")
}

// hashKeys 制作多个原始短链接key
func hashKeys(maxId uint64, num uint64) ([]eshort.Eshort, []interface{}) {
	arr := []eshort.Eshort{}
	var strs = []interface{}{}
	//不在此使用布隆过滤器，因为发生冲突的概率较小。
	//使用过滤器，通信频率太高，可能会造成redis阻塞。使用piple传输数量可能根据num，造成数量通信数据过大
	//故直接利用mysql唯一索引插入。
	ext := biz.BIZ.GetExt()
	for i := 1; i < int(num+1); i++ {
		skey := ehash.HashKey(maxId + uint64(i))
		arr = append(arr, eshort.Eshort{ShortKey: skey, Ext: ext})
		strs = append(strs, skey)
	}
	return arr, strs
}

func getMaxId() uint64 {
	type t struct {
		Mid sql.NullInt64
	}
	tt := t{}
	model.DB.Raw("SELECT MAX(id) as mid FROM eshorts").Scan(&tt)
	return cast.ToUint64(tt.Mid.Int64)
}

func inProduction() bool {
	bg := context.Background()
	i, err := eredis.RedisClient.Get(bg, produceStatus).Int()
	if err != nil {
		return false
	}
	if i == 1 {
		return true
	}
	return false
}

func Tidy() {
	var arr = []eshort.Eshort{}
	fmt.Println("进入到Tidy 函数")
	model.DB.Model(&eshort.Eshort{}).Where("status=0").Find(&arr)
	if len(arr) == 0 {
		fmt.Println("Tidy 查不到数据，进入生产")
		go BatchProductionSkey(0)
		return
	}
	fmt.Println("进入到Tidy 函数", "len:", len(arr))
	skeys := []interface{}{}
	strkeys := []string{}
	for _, shortUrl := range arr {
		skeys = append(skeys, shortUrl.ShortKey)
		strkeys = append(strkeys, shortUrl.ShortKey)
	}
	bg := context.Background()
	strings := eredis.RedisClient.LRange(bg, SkeyContainerName, 0, -1).Val()
	if len(strings) == 0 {
		err := eredis.RedisClient.RPush(context.Background(), SkeyContainerName, skeys...).Err()
		if err != nil {
			log.Fatal("数据库有未更新到缓存中，而缓存未更新成功")
		}
	}
	diff := array_tool.StrDiff(strkeys, strings)
	diffLen := len(diff)
	fmt.Println("diffLen", diffLen)
	if diffLen > 0 {
		dkeys := []interface{}{}
		for _, s := range diff {
			dkeys = append(dkeys, s)
		}
		fmt.Println("dkeys", dkeys)
		err := eredis.RedisClient.RPush(context.Background(), SkeyContainerName, dkeys...).Err()
		if err != nil {
			log.Fatal("数据库有未更新到缓存中，而缓存未更新成功")
		}
	}

}
