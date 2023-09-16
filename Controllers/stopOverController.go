package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CreateStopover(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	var body struct {
		AirportID uint
		Arrival   string
		Departure string
	}

	err = c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	arrivalTime, _ := time.Parse(time.RFC3339, body.Arrival)
	departureTime, _ := time.Parse(time.RFC3339, body.Departure)

	stopover := models.Stopover{
		FlightID:  uint(flightID),
		AirportID: body.AirportID,
		Arrival:   arrivalTime,
		Departure: departureTime,
	}

	result := initializers.DB.Create(&stopover)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{"stopover": stopover})
}

func GetAllStopovers(c *gin.Context) {
	flightIDStr := c.Param("id")
	fmt.Println(flightIDStr)
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	var stopovers []models.Stopover

	result := initializers.DB.Where("flight_id = ?", flightID).Find(&stopovers)
	fmt.Println(result, "result%%%%%%%%%%%%%%%%%%%%%%%%")
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"stopovers": stopovers})
}

func GetStopover(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	stopoverID := c.Param("stopID")

	var stopover models.Stopover
	result := initializers.DB.Where("flight_id = ?", flightID).First(&stopover, stopoverID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	c.JSON(200, gin.H{"stopover": stopover})
}

func UpdateStopover(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	stopoverID := c.Param("stopID")
	var body struct {
		AirportID uint
		Arrival   string
		Departure string
	}
	err = c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	var stopover models.Stopover
	result := initializers.DB.Where("flight_id = ?", flightID).First(&stopover, stopoverID)
	if result.Error != nil {
		c.Status(404)
		return
	}

	arrivalTime, _ := time.Parse(time.RFC3339, body.Arrival)
	departureTime, _ := time.Parse(time.RFC3339, body.Departure)

	stopover.AirportID = body.AirportID
	stopover.Arrival = arrivalTime
	stopover.Departure = departureTime

	result = initializers.DB.Save(&stopover)
	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{"stopover": stopover})
}

func DeleteStopover(c *gin.Context) {
	flightIDStr := c.Param("id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	stopoverID := c.Param("stopID")

	var stopover models.Stopover
	result := initializers.DB.Where("flight_id = ?", flightID).First(&stopover, stopoverID)

	if result.Error != nil {
		c.Status(404)
		return
	}

	result = initializers.DB.Delete(&stopover)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
