package api

import (
	"encoding/json"
	"errors"
	"fmt"
	mydb "github.com/StepanShevelev/fixedtask/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func apiCreateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}

	var user *mydb.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("an error occurred while decoding json"))
		return
	}
	//key := "user" + strconv.Itoa(user.ID)

	mydb.CreateUser(user)
	w.Write([]byte(strconv.Itoa(user.ID)))
	Caching.SetUserCache(user.ID, user)

	result, err := Caching.GetUserCache(user.ID)
	if err != nil {
		w.Write([]byte("an error occurred while getting cache"))
		w.Write([]byte(err.Error()))
		return
	}
	SendData(&result, w)
}

func apiGetUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}

	userId, okId := ParseId(w, r)
	if !okId {
		w.Write([]byte(`{"error": "can't pars id"}`))
		return
	}
	fmt.Println(userId)
	//key := "user" + strconv.Itoa(userId)

	result, err := Caching.GetUserCache(userId)
	//fmt.Sprint(result)
	if err != nil {
		w.Write([]byte("an error occurred while getting cache"))
		w.Write([]byte(err.Error()))
		return
	}
	SendData(&result, w)

	//
	//user, okUser := GetUserById(userId, w)
	//if !okUser {
	//	return
	//}
	//SendData(userId, w)
}

func GetUserById(userId int, w http.ResponseWriter) (*mydb.User, bool) {

	user, err := mydb.FindUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return user, true
}

func apiDeleteUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := ParseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := GetUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Database.Db.Unscoped().Delete(&usr)
	w.WriteHeader(http.StatusOK)
}

func apiUpdateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := ParseId(w, r)
	if !okId {
		return
	}

	user, okUsr := GetUserById(id, w)
	if !okUsr {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	mydb.Database.Db.Save(&user)

}
func apiGetAllUsers(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	users, okUsers := getUsers(w)
	if !okUsers {
		return
	}
	SendData(users, w)
}

func getUsers(w http.ResponseWriter) ([]mydb.User, bool) {

	users, err := mydb.FindUsers()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return users, true
}

func apiPharaohUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := ParseId(w, r)
	if !okId {
		return
	}
	var massivslice []mydb.UserCategories
	usr, okUsr := GetUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Database.Db.Where("user_id = ?", usr.ID).Unscoped().Delete(&massivslice)
	mydb.Database.Db.Where("user_id = ?", usr.ID).Unscoped().Delete(&mydb.Pet{})
	apiDeleteUser(w, r)
	w.WriteHeader(http.StatusOK)
}
