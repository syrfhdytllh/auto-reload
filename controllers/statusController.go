package controllers

import (
	"auto-reload/database"
	"auto-reload/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStatus(c *gin.Context) {
	var db = database.GetDB()

	var status models.Status
	//fmt.Println(status, "who are you")

	err := db.First(&status, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		statusinput := models.Status{Water: status.Water, Air: status.Air}
		db.Create(&statusinput)
		//fmt.Println(statusinput, "we are here")
		return
	}

	var input models.Status
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&status).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"data":    status,
		"message": "Update data success",
		"success": true})
}
