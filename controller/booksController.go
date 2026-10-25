package controller

import (
	"golang_quiz3/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// konversi ketebalan buku berdasarkan total page
func convertThickness(totalPage int) string {
	if totalPage > 100 {
		return "tebal"
	}
	return "tipis"
}

func CreateBook(c *gin.Context) {
	var book database.Books
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(book.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// konversi kolom thickness berdasarkan total_page
	book.Thickness = convertThickness(book.TotalPage)

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
		return
	}

	if book.CategoryID > 0 {
		var category database.Category
		if err := database.DB.First(&category, book.CategoryID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category_id does not exist"})
			return
		}
	}

	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(int); ok {
			book.ModifiedBy = id
		}
	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetAllBooks(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
		return
	}

	var books []database.Books
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
		return
	}

	id := c.Param("id")
	var book database.Books
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	var input database.Books
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ReleaseYear != 0 {
		if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
			return
		}
		book.ReleaseYear = input.ReleaseYear
	}

	if input.CategoryID > 0 && input.CategoryID != book.CategoryID {
		var category database.Category
		if err := database.DB.First(&category, input.CategoryID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category_id does not exist"})
			return
		}
		book.CategoryID = input.CategoryID
	}

	if strings.TrimSpace(input.Title) != "" {
		book.Title = input.Title
	}
	if strings.TrimSpace(input.Description) != "" {
		book.Description = input.Description
	}
	if input.ImageUrl != "" {
		book.ImageUrl = input.ImageUrl
	}
	if input.Price != 0 {
		book.Price = input.Price
	}
	if input.TotalPage != 0 {
		book.TotalPage = input.TotalPage
		book.Thickness = convertThickness(book.TotalPage)
	}

	if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(int); ok {
			book.ModifiedBy = uid
		}
	} else if input.ModifiedBy != 0 {
		book.ModifiedBy = input.ModifiedBy
	}

	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
		return
	}

	id := c.Param("id")
	var book database.Books
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func GetAllBookByCategory(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
		return
	}

	categoryID := c.Param("id")
	var category database.Category
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	var books []database.Books
	if err := database.DB.Where("category_id = ?", categoryID).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
