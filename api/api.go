package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InitBackendApi() {
	http.HandleFunc("/API/create_user", apiCreateUser)
	http.HandleFunc("/API/create_category", apiCreateCategory)
	http.HandleFunc("/API/create_pet", apiCreatePet)

	http.HandleFunc("/API/update_user", apiUpdateUser)
	http.HandleFunc("/API/update_category", apiUpdateCategory)
	http.HandleFunc("/API/update_pet", apiUpdatePet)

	http.HandleFunc("/API/get_user", apiGetUser)
	http.HandleFunc("/API/get_category", apiGetCategory)
	http.HandleFunc("/API/get_pet", apiGetPet)

	http.HandleFunc("/API/get_users", apiGetAllUsers)
	http.HandleFunc("/API/get_categories", apiGetAllCategories)
	http.HandleFunc("/API/get_pets", apiGetAllPets)

	http.HandleFunc("/API/delete_user", apiDeleteUser)
	http.HandleFunc("/API/delete_category", apiDeleteCategory)
	http.HandleFunc("/API/delete_pet", apiDeletePet)

	http.HandleFunc("/API/user_add_category", apiUserAddCategory)
	http.HandleFunc("/API/user_delete_category", apiUserDeleteCategory)

	http.HandleFunc("/API/pharaoh_user", apiPharaohUser)
}

func ParseId(w http.ResponseWriter, r *http.Request) (int, bool) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "arguments params are missing"}`))
		return 0, false
	}
	userId, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "can't pars id"}`))
		return 0, false
	}
	return userId, true
}

func SendData(data interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "can't marshal json"}`))
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func isMethodGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodDELETE(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodPOST(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodPUT(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
