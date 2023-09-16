package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreatePassenger(c *gin.Context) {
	var body struct {
		Name string
	}

	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	fmt.Println(" #######################", body)

	passenger := models.Passenger{Name: body.Name}

	result := initializers.DB.Create(&passenger)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{"passenger": passenger})
}

func GetAllPassengers(c *gin.Context) {
	var passengers []models.Passenger

	result := initializers.DB.Find(&passengers)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"passengers": passengers})
}

func GetPassenger(c *gin.Context) {
	passengerID := c.Param("id")

	var passenger models.Passenger
	result := initializers.DB.First(&passenger, passengerID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	c.JSON(200, gin.H{"passenger": passenger})
}

func UpdatePassenger(c *gin.Context) {
	passengerID := c.Param("id")
	var body struct {
		Name string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	var passenger models.Passenger
	result := initializers.DB.First(&passenger, passengerID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	// Update passenger's name only if it's not an empty string
	if body.Name != "" {
		passenger.Name = body.Name
		result = initializers.DB.Save(&passenger)
		if result.Error != nil {
			c.Status(500)
			return
		}
	}

	c.JSON(200, gin.H{"passenger": passenger})
}

func DeletePassenger(c *gin.Context) {
	passengerID := c.Param("id")

	var passenger models.Passenger
	result := initializers.DB.First(&passenger, passengerID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	result = initializers.DB.Delete(&passenger)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
