package database

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username   string    `gorm:"not null;unique" json:"username"`
	Password   string    `gorm:"not null" json:"password,omitempty"`
	ModifiedAt time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy int       `json:"modified_by"`
}

func (User) TableName() string { return "users" }

type Category struct {
	gorm.Model
	Name       string    `gorm:"not null" json:"name"`
	ModifiedAt time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy int       `json:"modified_by"`
}

func (Category) TableName() string { return "categories" }

type Books struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `gorm:"not null" json:"release_year"`
	Price       float64   `gorm:"not null" json:"price"`
	TotalPage   int       `gorm:"not null" json:"total_page"`
	Thickness   string    `gorm:"not null" json:"thickness"`
	CategoryID  int       `json:"category_id"`
	ModifiedAt  time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy  int       `json:"modified_by"`
}

func (Books) TableName() string { return "books" }
