package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.ruoyi.com/src/internal/pojo"
	"log"
	"time"
)

var redisClients *redis.Client

type NewRedis struct {
	client redis.Client
}
type Redistor interface {
	SetRedis(key string, value []byte, t time.Duration) error
	ExistRedis(key string) bool
	GetRedis(key string) string
	GetLimitRedis(key string) int
	DelRedis(key string) error
	ExpireRedis(key string, t time.Duration) error
	AutoInc(key string) error
	RpushRedis(name string, key string) error
}

func NewRedisClient() *NewRedis {
	return &NewRedis{client: *redisClients}
}

/**
 * @ClassName redis
 * @Description TODO
 * @Author khr
 * @Date 2023/7/31 11:02
 * @Version 1.0
 */
func Redisinit(redisClient *pojo.RedisConf) {
	redisClients = redis.NewClient(&redis.Options{
		Addr: redisClient.Host + ":" + redisClient.Port,
		//Username:   redisCon.UserName,
		//Password:   redisCon.PassWord,
		DB:         redisClient.Db,
		PoolSize:   redisClient.PoolSize,
		MaxRetries: redisClient.MaxRetries,
	})
	_, err := redisClients.Ping(context.Background()).Result()
	if err != nil {
		//fmt.Printf("redis connect get failed.%v", err.Error())
		log.Fatalf("redis connect get failed.%v", err.Error())
		return
	}
	log.Printf("redis 初始化连接成功")
}

// 添加数据
func (r *NewRedis) SetRedis(key string, value []byte, t time.Duration) error {
	return r.client.Set(context.Background(), key, value, t).Err()
	//expire := time.Duration(t)
	//if err = RedisClient.Set(ctx, key, value, t).Err(); err != nil {
	//	fmt.Println(err, "redis存放错误")
	//	return false
	//}
	//fmt.Println("redis存放时间", t)
	//return true
}

// set 中是否存在某个成员
func (r *NewRedis) ExistRedis(key string) bool {
	result, err := r.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// 获取数据
func (r *NewRedis) GetRedis(key string) string {
	result, _ := r.client.Get(context.Background(), key).Result()
	return result
}

// 获取数据
func (r *NewRedis) GetLimitRedis(key string) int {
	result, _ := r.client.Get(context.Background(), key).Int()
	return result
}

// 删除数据
func (r *NewRedis) DelRedis(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	return err
}

// 延长过期时间
func (r *NewRedis) ExpireRedis(key string, t time.Duration) error {
	if err := r.client.Expire(context.Background(), key, t).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
 * @MethodName
 * @Description redis自增
 * @Author khr
 * @Date 2023/7/31 15:25
 */
func (r *NewRedis) AutoInc(key string) error {
	return r.client.Incr(context.Background(), key).Err()
}

/*
 * @MethodName
 * @Description
 * @Author khr
 * @Date 2023/7/31 16:21
 */
func (r *NewRedis) RpushRedis(name string, key string) error {
	return r.client.RPush(context.Background(), name, key).Err()
}
