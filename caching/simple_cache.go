package caching

import (
	"github.com/h4lim/go-sdk/utils"
	"github.com/patrickmn/go-cache"
	"time"
)

var GlobalCache *cache.Cache

type SimpleCache struct {
	ExpiredTime int
	PurgeTime   int
}

type CacheFunction interface {
	CreateNewCache() *cache.Cache
	SetValue(key string, value string, cache cache.Cache)
	GetValue(key string, cache cache.Cache) *interface{}
	DeleteCache(key string, globalCache cache.Cache)
	SetValueWithHash(key string, value string, cache cache.Cache)
	GetValueWithHash(key string, cache cache.Cache) *interface{}
	DeleteCacheWithHash(key string, globalCache cache.Cache)

}

func (c *SimpleCache) CreateNewCache() *cache.Cache {

	newCache := cache.New(time.Minute*time.Duration(c.ExpiredTime),
		time.Minute*time.Duration(c.PurgeTime))
	return newCache

}

func (c *SimpleCache) SetValue(key string, value string, globalCache cache.Cache) {

	globalCache.Set(key, value, time.Minute*time.Duration(c.ExpiredTime))

}

func (c *SimpleCache) GetValue(key string, globalCache cache.Cache) *interface{} {

	data, found := globalCache.Get(key)
	if found {
		return &data
	}

	return nil

}

func (c *SimpleCache) DeleteCache(key string, globalCache cache.Cache) {

	globalCache.Delete(key)

}

func (c *SimpleCache) SetValueWithHash(key string, value string, globalCache cache.Cache) {

	globalCache.Set(utils.EncryptionSha256(key), utils.EncryptionSha256(value),
		time.Minute*time.Duration(c.ExpiredTime))

}

func (c *SimpleCache) GetValueWithHash(key string, globalCache cache.Cache) *interface{} {

	data, found := globalCache.Get(utils.EncryptionSha256(key))
	if found {
		return &data
	}

	return nil

}

func (c *SimpleCache) DeleteCacheWithHash(key string, globalCache cache.Cache) {

	globalCache.Delete(utils.EncryptionSha256(key))

}
