package handler

import (
	"net/http"
	"os"
	"strings"

	"golang_quiz3/database"
	"golang_quiz3/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var app http.Handler

func init() {
	psqlInfo := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if psqlInfo != "" {
		var err error
		database.DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err == nil {
			_ = database.DB.AutoMigrate(&database.User{}, &database.Category{}, &database.Books{})
		}
	}
	app = routers.StartServer()
}

func Handler(w http.ResponseWriter, r *http.Request) {
    // menjalankan engine Gin
	app.ServeHTTP(w, r)
}
