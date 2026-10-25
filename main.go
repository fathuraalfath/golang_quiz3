package main

import (
	"fmt"
	"golang_quiz3/database"
	"golang_quiz3/routers"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
		sslmode := getEnv("DB_SSLMODE", "disable")

		fmt.Printf("Mencoba koneksi ke database host: %s, port: %s, dbname: %s\n", host, port, dbname)

		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)
	}

	// Buka koneksi ke database menggunakan GORM
	var err error
	database.DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Printf("Peringatan: Gagal terkoneksi ke database (%v)\n", err)
	} else {
		fmt.Println("Database berhasil terkoneksi")
		_ = database.DB.AutoMigrate(&database.User{}, &database.Category{}, &database.Books{})
	}

	// Menjalankan server Gin dengan port dinamis dari environment
	portServer := getEnv("PORT", "8080")
	routers.StartServer().Run(":" + portServer)
}
