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
	  // 1. Ambil path asli yang diminta dari Postman/browser
    originalPath := r.URL.Path

    // 2. Jika path tidak diawali dengan /api, tambahkan /api di depannya
    if !strings.HasPrefix(originalPath, "/api") {
        if originalPath == "/" {
            r.URL.Path = "/api"
        } else {
            r.URL.Path = "/api" + originalPath
        }
    } else {
        r.URL.Path = originalPath
    }

    // 3. Jalankan aplikasi Gin 
    app.ServeHTTP(w, r)
}
