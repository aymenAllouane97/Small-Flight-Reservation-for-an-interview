package controllers

import (
	initializers "awesomeProject/Initializers" // Import your DB initialization here
	models "awesomeProject/Models"             // Import your City model here
	"github.com/gin-gonic/gin"
)

func CreateCity(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	city := models.City{Name: body.Name}

	result := initializers.DB.Create(&city)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{"city": city})
}

func GetAllCities(c *gin.Context) {
	var cities []models.City

	result := initializers.DB.Find(&cities)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"cities": cities})
}

func GetCity(c *gin.Context) {
	cityID := c.Param("id")

	var city models.City
	result := initializers.DB.First(&city, cityID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	c.JSON(200, gin.H{"city": city})
}

func UpdateCity(c *gin.Context) {
	cityID := c.Param("id")

	var body struct {
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var city models.City
	result := initializers.DB.First(&city, cityID)
	if result.Error != nil {
		c.Status(404)
		return
	}

	city.Name = body.Name
	result = initializers.DB.Save(&city)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"city": city})
}

func DeleteCity(c *gin.Context) {
	cityID := c.Param("id")

	var city models.City
	result := initializers.DB.First(&city, cityID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	result = initializers.DB.Delete(&city)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
