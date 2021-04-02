package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisInterface interface {
	Get(key string) (string, error)
	Set(key string, value string, TTL time.Duration) (string, error)
	Exists(key ...string) (int64, error)
	Exist(key string) (bool, error)
}
type redisClient interface {
	Get(key string) *redis.StringCmd
	HSet(key, field string, value interface{}) *redis.BoolCmd
	HGet(key, field string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Exists(keys ...string) *redis.IntCmd
}
type redisModule struct {
	client redisClient
}

func New(address, password string) RedisInterface {

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	return &redisModule{
		client: client,
	}

}
func (r *redisModule) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}
func (r *redisModule) HGet(key, field string) (string, error) {
	return r.client.HGet(key, field).Result()
}
func (r *redisModule) Set(key string, value string, TTL time.Duration) (string, error) {
	return r.client.Set(key, value, TTL).Result()
}
func (r *redisModule) HSet(key, field string, value string, TTL time.Duration) (bool, error) {
	return r.client.HSet(key, field, value).Result()
}
func (r *redisModule) Exists(key ...string) (int64, error) {
	return r.client.Exists(key...).Result()
}

func (r *redisModule) Exist(key string) (bool, error) {
	count, err := r.client.Exists(key).Result()
	if count > 0 {
		return true, err
	}
	return false, err
}
