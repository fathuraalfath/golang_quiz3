package controller_test

import (
	"bytes"
	"encoding/json"
	"golang_quiz3/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestAuthMiddlewareUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/category", nil)

	handler := controller.AuthMiddleware()
	handler(c)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401 Unauthorized without auth header, got %d", w.Code)
	}
}

func TestCreateBookReleaseYearValidation(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// test release year < 1980
	payload := map[string]interface{}{
		"title":        "Buku Lama",
		"release_year": 1975,
		"total_page":   50,
	}
	body, _ := json.Marshal(payload)
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	controller.CreateBook(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for release_year < 1980, got %d", w.Code)
	}
}
