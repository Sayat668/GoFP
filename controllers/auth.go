package controllers

import (
	// "fmt"
	// "fproj/models"
	// "fproj/utils"
	// "net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
	"fproj/models"
	"gorm.io/gorm"
)

// using db in pages
var db *gorm.DB

func SetDB(DB *gorm.DB) {
	db = DB
}

// handlers here
func RegisterHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Получаем данные из запроса
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Здесь вы можете добавить свою логику валидации, например, проверить, что поля не пустые.

	// Создаем нового пользователя
	newUser := models.User{
		Name:  name,
		Email: email,
		Password: password,
		// Установите начальные значения для других полей, если необходимо.
	}

	// Сохраняем нового пользователя в базу данных
	if err := db.Create(&newUser).Error; err != nil {
		// Обработка ошибок сохранения
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Отправляем клиенту успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Registered successfully!",
		"user":    newUser,
	})
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/main.html")
}
