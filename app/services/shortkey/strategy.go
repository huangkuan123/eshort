package shortkey

import (
	"context"
	"errors"
	"eshort/app/services/shortkey/factory"
	"eshort/pkg/config"
	"eshort/pkg/eredis"
	"github.com/go-redis/redis/v8"
	"time"
)

var growType string

func Setup() {
	growType = config.GetString("eshort.grow_type")
}

// 获取策略选择
func Take() (string, error) {
	switch growType {
	case "default":
		return quicklyMade()
	case "initiative":
		return getSkeyInitiative()
	case "passive":
		return getSkeyPassive()
	}
	return "", errors.New("策略选择有误")
}

// 获取单条skey
// 主动扩容，该方法使用了pipeline，不可redis集群。
// 主动判断容量是否需要补齐，从而生产多个shortkey
func quicklyMade() (string, error) {
	rc := eredis.RedisClient
	pipeline := rc.Pipeline()
	bg := context.Background()
	//集群
	pipeline.LPop(bg, factory.SkeyContainerName)
	pipeline.LLen(bg, factory.SkeyContainerName)
	res, err := pipeline.Exec(bg)
	if err != nil {
		return "", err
	}
	cmd, ok := res[1].(*redis.IntCmd)
	if !ok {
		return "", cmd.Err()
	}
	slen, err := cmd.Result()
	if err != nil {
		return "", err
	}
	idcmd, ok := res[0].(*redis.StringCmd)
	if !ok {
		return "", cmd.Err()
	}
	skey, err := idcmd.Result()
	if err != nil {
		return "", err
	}
	if len(skey) == 0 { //调用生成服务，立即获取一条。
		go factory.BatchProductionSkey(slen)
		return factory.MakeRowSkey()
	}
	Produce(slen)
	return skey, nil
}

// 主动扩容，可集群
func getSkeyInitiative() (string, error) {
	skey, slen, err := getSkeySlen()
	if err != nil {
		return "", err
	}
	if len(skey) == 0 { //调用生成服务，立即获取一条。
		go factory.BatchProductionSkey(slen)
		return factory.MakeRowSkey()
	}
	Produce(slen)
	return skey, nil
}

// 被动扩容，可集群
func getSkeyPassive() (string, error) {
	rc := eredis.RedisClient
	bg := context.Background()
	skey := rc.LPop(bg, factory.SkeyContainerName).Val()
	if len(skey) == 0 { //调用生成服务，立即获取一条。
		return factory.MakeRowSkey()
	}
	return skey, nil
}

func getSkeySlen() (string, int64, error) {
	rc := eredis.RedisClient
	bg := context.Background()
	skey := rc.LPop(bg, factory.SkeyContainerName).Val()
	ulen, err := rc.LLen(bg, factory.SkeyContainerName).Uint64()
	if err != nil {
		return "", 0, err
	}
	slen := int64(ulen)
	return skey, slen, nil
}

//监控
func SuperPool() {
	if growType == "passive" {
		rc := eredis.RedisClient
		bg := context.Background()
		go func() {
			for {
				time.Sleep(60 * time.Second)
				lLen, _ := rc.LLen(bg, factory.SkeyContainerName).Uint64()
				Produce(int64(lLen))
			}
		}()
	}
}

func Produce(slen int64) {
	if factory.TooLittle(slen) { //低于阈值，就去生产
		go factory.BatchProductionSkey(slen)
	}
}
