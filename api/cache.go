package api

import (
	"github.com/bluele/gcache"
	"time"
)

type Cache struct {
	UserCache     gcache.Cache
	CategoryCache gcache.Cache
	NewCache      gcache.Cache
}

var Caching Cache

func (c *Cache) CreteCache() gcache.Cache {

	gc := gcache.New(50).
		LRU().
		Build()

	//Caching = Cache{
	//	UserCache: gc,
	//}
	c.UserCache = gc

	return gc
}
func (c *Cache) CreteCateCache() gcache.Cache {

	gc := gcache.New(50).
		LRU().
		Build()

	//Caching = Cache{
	//	CategoryCache: gc,
	//}
	c.CategoryCache = gc

	return gc
}

func (c *Cache) SetUserCache(key interface{}, value interface{}) {
	c.CreteCache()

	err := Caching.UserCache.SetWithExpire(key, value, time.Minute*5)
	if err != nil {
		return
	}
}

func (c *Cache) GetUserCache(key interface{}) (interface{}, error) {
	get, err := Caching.UserCache.Get(key)
	if err != nil {
		return "", err
	}

	return get, nil
}

func (c *Cache) SetCategoryCache(key interface{}, value interface{}) {
	c.CreteCateCache()

	err := Caching.CategoryCache.SetWithExpire(key, value, time.Minute*5)
	if err != nil {
		return
	}
}

func (c *Cache) GetCategoryCache(key interface{}) (interface{}, error) {
	get, err := Caching.CategoryCache.Get(key)
	if err != nil {
		return "", err
	}

	return get, nil
}

//func (c *Cache) SetCache(key interface{}, value interface{}) {
//	c.CreteCache()
//
//	err := Caching.NewCache.SetWithExpire(key, value, time.Minute*5)
//	if err != nil {
//		return
//	}
//
//}
//
//func (c *Cache) GetCache(key interface{}) (interface{}, error) {
//	get, err := Caching.NewCache.Get(key)
//	if err != nil {
//		return "", err
//	}
//
//	return get, nil
//}
