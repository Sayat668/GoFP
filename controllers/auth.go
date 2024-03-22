package controllers

import (
	// "fmt"
	// "fproj/models"
	// "fproj/utils"
	// "net/http"
	// "strconv"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
//using db in pages
var db *gorm.DB
func SetDB(DB *gorm.DB) {
	db = DB
}

//handlers here
