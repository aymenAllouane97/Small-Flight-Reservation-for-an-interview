package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateFlight(c *gin.Context) {
	var body struct {
		CompanyID          uint
		DepartureAirportID uint
		ArrivalAirportID   uint
		DepartureDate      string
		ArrivalDate        string
		DepartureHour      string
		ArrivalHour        string
	}

	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	fmt.Println(" #######################", body)

	departureDate, _ := time.Parse("02-01-2006", body.DepartureDate)
	arrivalDate, _ := time.Parse("02-01-2006", body.ArrivalDate)
	departureHour, _ := time.Parse("15:04:05", body.DepartureHour)
	arrivalHour, _ := time.Parse("15:04:05", body.ArrivalHour)

	flight := models.Flight{
		CompanyID:          body.CompanyID,
		DepartureAirportID: body.DepartureAirportID,
		ArrivalAirportID:   body.ArrivalAirportID,
		DepartureDate:      departureDate,
		ArrivalDate:        arrivalDate,
		DepartureHour:      departureHour,
		ArrivalHour:        arrivalHour,
	}

	result := initializers.DB.Create(&flight)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{"flight": flight})
}

func GetAllFlights(c *gin.Context) {
	var flights []models.Flight

	result := initializers.DB.Preload("Airports").Find(&flights)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"flights": flights})
}

func GetFlight(c *gin.Context) {
	flightID := c.Param("id")

	var flight models.Flight
	result := initializers.DB.Preload("Airports").First(&flight, flightID)

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{"flight": flight})
}

func UpdateFlight(c *gin.Context) {
	flightID := c.Param("id")
	var body struct {
		CompanyID          uint
		DepartureAirportID uint
		ArrivalAirportID   uint
		DepartureDate      string
		ArrivalDate        string
		DepartureHour      string
		ArrivalHour        string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	var flight models.Flight
	result := initializers.DB.First(&flight, flightID)
	if result.Error != nil {
		c.Status(404)
		return
	}
	departureDate, _ := time.Parse("2006-01-02", body.DepartureDate)
	arrivalDate, _ := time.Parse("2006-01-02", body.ArrivalDate)
	departureHour, _ := time.Parse("15:04:05", body.DepartureHour)
	arrivalHour, _ := time.Parse("15:04:05", body.ArrivalHour)

	flight.CompanyID = body.CompanyID
	flight.DepartureAirportID = body.DepartureAirportID
	flight.ArrivalAirportID = body.ArrivalAirportID
	flight.DepartureDate = departureDate
	flight.ArrivalDate = arrivalDate
	flight.DepartureHour = departureHour
	flight.ArrivalHour = arrivalHour

	result = initializers.DB.Save(&flight)
	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{"flight": flight})
}

func DeleteFlight(c *gin.Context) {
	flightID := c.Param("id")

	var flight models.Flight
	result := initializers.DB.First(&flight, flightID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	result = initializers.DB.Delete(&flight)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
