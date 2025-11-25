package controllers

import (
	config "gorm/Config"
	models "gorm/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	var count int64
	// Get actual user count from DB
	if err := config.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{
			"Title": "Admin Dashboard",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Title":     "Admin Dashboard",
		"UserCount": count, // use actual count from DB
	})
}

func Users(c *gin.Context) {
	var users []models.User

	if err := config.DB.Order("id desc").Find(&users).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "users.html", gin.H{
			"Title": "Users",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Title":     "Admin Dashboard",
		"UserCount": 2, // use actual count from DB
	})
}
