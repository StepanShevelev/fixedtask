package db

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {
	dsn := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable timezone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Pet{})
	db.AutoMigrate(&Category{})

	Database = DbInstance{
		Db: db,
	}
}

func CreateUser(r *http.Request) {
	var user *User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	Database.Db.Create(&user)
}

func FindUserById(Id int) (*User, error) {
	var user *User

	result := Database.Db.Preload("Pets").Preload("Categories").Find(&user, "id = ?", Id)

	if result.Error != nil {
		return nil, nil
	}
	return user, nil
}

func FindUsers() ([]User, error) {
	var users []User

	result := Database.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func CreateCategory(r *http.Request) {
	var category *Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		return
	}
	Database.Db.Create(&category)
}

func AddCategory(r *http.Request) {

	type UserCategoriesS struct {
		UserId      int   `json:"user_id"`
		CategoryIds []int `json:"category_ids"`
	}
	var ids *UserCategoriesS
	var massivslice []UserCategories

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		return
	}

	for _, categoryId := range ids.CategoryIds {

		massivslice = append(massivslice, UserCategories{UserId: ids.UserId, CategoryId: categoryId})

	}
	Database.Db.Create(&massivslice)
}

func DeleteCategory(r *http.Request) {

	type UserCategoriesS struct {
		UserId      int   `json:"user_id"`
		CategoryIds []int `json:"category_ids"`
	}
	var ids *UserCategoriesS
	var massivslice []UserCategories

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		return
	}

	for _, categoryId := range ids.CategoryIds {

		massivslice = append(massivslice, UserCategories{UserId: ids.UserId, CategoryId: categoryId})

	}
	Database.Db.Where("user_id = ?", ids.UserId).Where("category_id = ?", ids.CategoryIds).Unscoped().Delete(&massivslice)
}

func FindCategoryById(Id int) (*Category, error) {
	var category *Category
	result := Database.Db.Find(&category, "id = ?", Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
func FindCategories() ([]Category, error) {
	var categories []Category

	result := Database.Db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func CreatePet(r *http.Request) {
	var pet *Pet
	var user *User
	err := json.NewDecoder(r.Body).Decode(&pet)
	Database.Db.Model(&pet).Association("pets").Append(&user)
	if err != nil {
		return
	}
	Database.Db.Create(&pet)
}

func FindPetById(Id int) (*Pet, error) {
	var pet *Pet
	result := Database.Db.Find(&pet, "id = ?", Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return pet, nil
}

func FindUserPetByID(Id int) (*Pet, error) {
	var pet *Pet
	result := Database.Db.Find(&pet, "user_id = ?", Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return pet, nil
}

func FindPets() ([]Pet, error) {
	var pets []Pet

	result := Database.Db.Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}
	return pets, nil
}
