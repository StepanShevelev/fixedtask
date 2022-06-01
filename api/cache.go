package api

import (
	"fmt"
	mydb "github.com/StepanShevelev/fixedtask/db"
	"github.com/bluele/gcache"
	"net/http"
	"time"
)

func CreateUserCache(w http.ResponseWriter, r *http.Request, user *mydb.User) {

	gc := gcache.New(20).
		Simple().
		Build()
	err := gc.SetWithExpire(&user.Name, &user, time.Minute*5)
	if err != nil {
		w.Write([]byte("an error occurred while caching user data"))
		return
	}

	w.Write([]byte("success, user data has been cached"))

	return

}

func LoadUserCache(w http.ResponseWriter, r *http.Request) {

	userId, ok := ParseId(w, r)
	if !ok {
		w.Write([]byte("an error occurred while parsing id"))
		return
	}

	var evictCounter, loaderCounter, purgeCounter int
	gc := gcache.New(20).
		Simple().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			loaderCounter++
			expire := 5 * time.Minute
			return "expired", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			evictCounter++
			fmt.Println("evicted key:", key)
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			purgeCounter++
			fmt.Println("purged key:", key)
		}).
		Build()
	value, err := gc.Get(&userId)
	if err != nil {
		w.Write([]byte("an error occurred while searching user data in cache"))
		return
	}

	if loaderCounter != evictCounter+purgeCounter {
		w.Write([]byte("user not found"))

	}

	SendData(value, w)
	return
}

func CreatePetCache(w http.ResponseWriter, r *http.Request, pet *mydb.Pet) {

	gc := gcache.New(20).
		Simple().
		Build()
	err := gc.SetWithExpire(&pet.Name, &pet, time.Minute*5)
	if err != nil {
		w.Write([]byte("an error occurred while caching user data"))
		return
	}

	w.Write([]byte("success, user data has been cached"))

	return

}

func LoadPetCache(w http.ResponseWriter, r *http.Request) {

	userId, ok := ParseId(w, r)
	if !ok {
		w.Write([]byte("an error occurred while parsing id"))
		return
	}

	var evictCounter, loaderCounter, purgeCounter int
	gc := gcache.New(20).
		Simple().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			loaderCounter++
			expire := 5 * time.Minute
			return "expired", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			evictCounter++
			fmt.Println("evicted key:", key)
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			purgeCounter++
			fmt.Println("purged key:", key)
		}).
		Build()
	value, err := gc.Get(&userId)
	if err != nil {
		w.Write([]byte("an error occurred while searching user data in cache"))
		return
	}

	if loaderCounter != evictCounter+purgeCounter {

		w.Write([]byte("user not found"))

	}

	SendData(value, w)
	return
}