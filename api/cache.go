package api

import (
	"github.com/bluele/gcache"
	"time"
)

type Cache struct {
	NewCache gcache.Cache
}

var Caching Cache

func (c *Cache) CreteCache() gcache.Cache {

	gc := gcache.New(50).
		LRU().
		Build()

	Caching = Cache{
		NewCache: gc,
	}

	return gc
}

func (c *Cache) SetCache(key interface{}, value interface{}) {
	c.CreteCache()

	err := Caching.NewCache.SetWithExpire(key, value, time.Minute*5)
	if err != nil {
		return
	}

}

func (c *Cache) GetCache(key interface{}) (interface{}, error) {
	get, err := Caching.NewCache.Get(key)
	if err != nil {
		return "", err
	}

	return get, nil
}

//func CreateUserCache(w http.ResponseWriter, user *mydb.User) {
//	fmt.Println(user.Name)
//	fmt.Println(user.ID)
//
//	gc := gcache.New(50).
//		LRU().
//		Build()
//	err := gc.Set(user.ID, user)
//	value, err := gc.Get(user.ID)
//	fmt.Println("Get:", value)
//	fmt.Println(err)
//	if err != nil {
//		w.Write([]byte("an error occurred while caching user data"))
//		return
//	}
//	w.Write([]byte("success, user data has been cached"))
//	SendData(value, w)
//	return
//
//}
//
//func LoadUserCache(w http.ResponseWriter, r *http.Request) {
//
//	userId, ok := ParseId(w, r)
//	if !ok {
//		w.Write([]byte("an error occurred while parsing id"))
//		return
//	}
//
//	//var evictCounter, purgeCounter int
//	gc := gcache.New(20).
//		LRU().
//		//LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
//		//	loaderCounter++
//		//	expire := 5 * time.Minute
//		//	return "expired", &expire, nil
//		//}).
//		//EvictedFunc(func(key, value interface{}) {
//		//	evictCounter++
//		//	fmt.Println("evicted key:", key)
//		//}).
//		//PurgeVisitorFunc(func(key, value interface{}) {
//		//	purgeCounter++
//		//	fmt.Println("purged key:", key)
//		//}).
//		Build()
//	fmt.Println(userId)
//	value, _ := gc.Get(userId)
//	fmt.Println(value)
//	//if err == nil {
//	//	w.Write([]byte("an error occurred while searching user data in cache"))
//	//	return
//	//}
//
//	//if loaderCounter != evictCounter+purgeCounter
//	if value == nil {
//
//		w.Write([]byte("user not found"))
//
//	}
//
//	SendData(value, w)
//	return
//}
