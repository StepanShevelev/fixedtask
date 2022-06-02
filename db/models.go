package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         int `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *gorm.DeletedAt `gorm:"index"`
	Name       string          `json:"name" db:"name"`
	Categories []Category      `json:"categories" db:"categories" gorm:"many2many:user_categories;"`
	Pets       []Pet           `json:"pets" db:"pets" gorm:"foreignKey:UserID"`
}

type UserCategories struct {
	UserId     int `json:"user_id"`
	CategoryId int `json:"category_id"`
}

type Category struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	Users     []User          `json:"users" db:"users" gorm:"many2many:user_categories;"`
	Name      string          `json:"name" db:"name"`
}

type Pet struct {
	gorm.Model
	Name    string `json:"name" db:"name"`
	Counter int    `json:"counter" db:"counter"`
	UserID  int    `json:"user_id" db:"user_id"`
}

// ErrLogs storage some error logs
type ErrLogs struct {
	gorm.Model
	Error string `json:"error" db:"error"`
	Place string `json:"place" db:"place"`
	Count int    `json:"count" db:"count"`
}
