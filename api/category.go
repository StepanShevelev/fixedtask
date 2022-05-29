package api

import (
	"encoding/json"
	"errors"
	mydb "github.com/StepanShevelev/fixedtask/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func apiCreateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}

	mydb.CreateCategory(r)

}

func apiUpdateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := parseId(w, r)
	if !okId {
		return
	}

	category, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		return
	}

	mydb.Database.Db.Save(&category)

}
func apiGetAllCategories(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	categories, okCats := getCategories(w)
	if !okCats {
		return
	}
	sendData(categories, w)
}

func getCategories(w http.ResponseWriter) ([]mydb.Category, bool) {

	categories, err := mydb.FindCategories()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return categories, true
}

func apiGetCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	categoryId, okId := parseId(w, r)
	if !okId {
		return
	}

	category, okCat := getCategoryById(categoryId, w)
	if !okCat {
		return
	}
	sendData(category, w)
}

func getCategoryById(categoryId int, w http.ResponseWriter) (*mydb.Category, bool) {

	category, err := mydb.FindCategoryById(categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return category, true
}

func apiDeleteCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	category, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	mydb.Database.Db.Delete(&category)
	w.WriteHeader(http.StatusOK)
}

func apiUserAddCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}
	mydb.AddCategory(r)
}

func apiUserDeleteCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}

	mydb.DeleteCategory(r)
}
