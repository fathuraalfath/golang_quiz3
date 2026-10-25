package main

import (
	"database/sql"
	"fmt"
	"golang_quiz3/routers"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	db  *sql.DB
	err error
)

// tabel users
type User struct {
	gorm.Model
	username    string    `gorm:"not null"`
	password    string    `gorm:"not null;unique"`
	modified_at time.Time `gorm:"autoUpdateTime"`
	modified_by int
}

// tabel buku
type Books struct {
	gorm.Model
	title        string `gorm:"not null"`
	description  string `gorm:"not null"`
	image_url    string
	release_year int     `gorm:"not null"`
	price        float64 `gorm:"not null"`
	total_page   int     `gorm:"not null"`
	thickness    int     `gorm:"not null"`
	category_id  int
	modified_at  time.Time `gorm:"autoUpdateTime"`
	modified_by  int       `gorm:"not null"`
}

// tabel kategori
type Category struct {
	gorm.Model
	name        string    `gorm:"not null"`
	modified_at time.Time `gorm:"autoUpdateTime"`
	modified_by int       `gorm:"not null"`
}

// func untuk membaca .env
func getEnv(key, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return fallback
}

// main function
func main() {
	_ = godotenv.Load()
	var psqlInfo string

	if dbURL := strings.TrimSpace(os.Getenv("DATABASE_URL")); dbURL != "" {
		psqlInfo = dbURL
		fmt.Println("Menggunakan konfigurasi dari DATABASE_URL")
	} else {
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "")
		dbname := getEnv("DB_NAME", "postgres")
		sslmode := getEnv("DB_SSLMODE", "require")

		fmt.Printf("Mencoba koneksi ke database host: %s, port: %s, dbname: %s\n", host, port, dbname)

		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)
	}

	// Buka koneksi ke database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifikasi koneksi ke database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Pesan sukses koneksi
	fmt.Println("Database berhasil terkoneksi")

	// Menjalankan server Gin dengan port dinamis dari environment
	portServer := getEnv("PORT", "8080")
	routers.StartServer().Run(":" + portServer)
}
