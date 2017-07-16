package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itheodoro/rest-mysql/db"
	"github.com/itheodoro/rest-mysql/models"
)

// CreateDog - Create a new dog
func CreateDog(c *gin.Context) {
	db := db.InitDB()
	defer db.Close()

	dog := models.Dog{}

	c.Bind(&dog)

	if res := db.Create(&dog); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.JSON(http.StatusCreated, gin.H{"record": dog})
	}

}

// GetAllDogs - Return all dogs
func GetAllDogs(c *gin.Context) {
	db := db.InitDB()
	defer db.Close()

	var dogs []models.Dog

	if res := db.Find(&dogs); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{"dogs": dogs})
	}

}
