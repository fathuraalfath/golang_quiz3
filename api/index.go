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
	 // 1. Tangkap rute asli utuh dari Postman melalui Header internal Vercel
    forwardedUri := r.Header.Get("X-Forwarded-Uri")
    matchedPath := r.Header.Get("X-Matched-Path")

    // 2. Timpa rute request
    if forwardedUri != "" {
        r.URL.Path = forwardedUri
    } else if matchedPath != "" {
        r.URL.Path = matchedPath
    }
    if idx := strings.Index(r.URL.Path, "?"); idx != -1 {
        r.URL.Path = r.URL.Path[:idx]
    }

    app.ServeHTTP(w, r)
}
