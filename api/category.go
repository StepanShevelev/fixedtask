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

func apiCreateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}

	var category *mydb.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.Write([]byte("an error occurred while decoding json"))
		return
	}
	key := "category" + strconv.Itoa(category.ID)
	mydb.CreateCategory(r, category)
	Caching.SetCache(key, category)

	//
	///
	///
	result, err := Caching.GetCache(key)
	if err != nil {
		w.Write([]byte("an error occurred while getting cache"))
		return
	}
	SendData(&result, w)

}

func apiUpdateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := ParseId(w, r)
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

	SendData(categories, w)
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
	//categoryId, okId := ParseId(w, r)
	//if !okId {
	//	return
	//}
	//
	//category, okCat := getCategoryById(categoryId, w)
	//if !okCat {
	//	return
	//}
	var category *mydb.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.Write([]byte("an error occurred while decoding json"))
		return
	}

	userHead := mydb.Database.Db.Find(&category, "name = ?", category.Name)
	if userHead.Error != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		//mydb.UppendErrorWithPath(userHeader.Error)
		return
	}

	//categoryId, okId := ParseId(w, r)
	//if !okId {
	//	return
	//}

	//LoadUserCache(w, r)
	result, err := Caching.GetCache(category.Name)
	fmt.Sprint(result)
	if err != nil {
		w.Write([]byte("an error occurred while getting cache"))
		return
	}
	SendData(result, w)
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
	id, okId := ParseId(w, r)
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

	type UserCategoriesS struct {
		Username     string `json:"username"`
		UserId       int    `json:"user_id"`
		Categoryname string `json:"categoryname"`
		CategoryId   int    `json:"category_id"`
	}
	var category *mydb.Category
	var user *mydb.User
	var ids UserCategoriesS

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		return
	}
	userHead := mydb.Database.Db.Find(&category, "id = ?", ids.CategoryId)
	if userHead.Error != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		//mydb.UppendErrorWithPath(userHeader.Error)
		return
	}

	userHeader := mydb.Database.Db.Find(&user, "id = ?", ids.UserId)
	if userHeader.Error != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		//mydb.UppendErrorWithPath(userHeader.Error)
		return
	}

	category.Users = append(category.Users, mydb.User{ID: ids.UserId})

	mydb.Database.Db.Model(&category).Association("Users").Append(&user)
	w.Write([]byte(strconv.Itoa(ids.UserId)))

	Caching.SetCache(ids.UserId, user)
	Caching.SetCache(ids.CategoryId, category)

	//var user mydb.User
	//var category mydb.Category
	//
	//
	//
	//mydb.Database.Db.Model(&category).Association("Users").Append(&user)
	//
	//
	//
	//

}

func apiUserDeleteCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}

	mydb.DeleteCategory(r)
}
