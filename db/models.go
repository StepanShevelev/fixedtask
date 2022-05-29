package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string     `json:"name" db:"name"`
	Categories []Category `json:"categories" db:"categories" gorm:"many2many:user_categories;"`
	Pets       []Pet      `json:"pets" db:"pets" gorm:"foreignKey:UserID"`
}

type UserCategories struct {
	UserId     int `json:"user_id"`
	CategoryId int `json:"category_id"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" db:"name"`
}

type Pet struct {
	gorm.Model
	Name    string `json:"name" db:"name"`
	Counter int    `json:"counter" db:"counter"`
	UserID  int    `json:"user_id" db:"user_id"`
}
