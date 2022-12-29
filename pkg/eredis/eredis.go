package eredis

import (
	"context"
	"eshort/pkg/config"
	"eshort/pkg/easylogger"
	"eshort/pkg/ehash"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host") + ":" + config.GetString("redis.port"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.select_db"), // use default DB
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		easylogger.LogError(err, "。redis连接失败")
	}
	return RedisClient
}

func JoinBloom(key string) {
	RedisClient.SetBit(context.Background(), "eshort:bloom", ehash.Decode(key), 1)
}

func InBloom(key string) bool {
	return RedisClient.GetBit(context.Background(), "eshort:bloom", ehash.Decode(key)).Val() == 1
}

func GetLock(key string, uuid int64, exp time.Duration) {
	timeout := exp * time.Second
	hasLock := false
	for !hasLock {
		hasLock = RedisClient.SetNX(context.Background(), key, uuid, timeout).Val()
		if !hasLock {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
}

func Unlock(key string, uuid int64) (bool, error) {
	script := "if redis.call('get',KEYS[1]) == ARGV[1] then \n  return redis.call('del',KEYS[1]) else\n  return 0 end"
	eval := RedisClient.Eval(context.Background(), script, []string{key}, uuid)
	res, err := eval.Int()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}
