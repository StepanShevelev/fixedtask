package api

import (
	"encoding/json"
	"errors"
	mydb "github.com/StepanShevelev/fixedtask/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func apiCreateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}
	mydb.CreateUser(r)
}

func apiUpdateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := parseId(w, r)
	if !okId {
		return
	}

	user, okUsr := getUserById(id, w)
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
	sendData(users, w)
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

func apiGetUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	userId, okId := parseId(w, r)
	if !okId {
		return
	}

	user, okUser := getUserById(userId, w)
	if !okUser {
		return
	}
	sendData(user, w)
}

func getUserById(userId int, w http.ResponseWriter) (*mydb.User, bool) {

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
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Database.Db.Unscoped().Delete(&usr)
	w.WriteHeader(http.StatusOK)
}

func apiPharaohUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}
	var massivslice []mydb.UserCategories
	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Database.Db.Where("user_id = ?", usr.ID).Unscoped().Delete(&massivslice)
	mydb.Database.Db.Where("user_id = ?", usr.ID).Unscoped().Delete(&mydb.Pet{})
	apiDeleteUser(w, r)
	w.WriteHeader(http.StatusOK)
}
