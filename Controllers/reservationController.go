package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateReservation(c *gin.Context) {
	clientIDStr := c.Param("id")
	clientID, err := strconv.ParseUint(clientIDStr, 10, 64)
	fmt.Println(clientID, "_______________________)))))))")
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	flightIDStr := c.Param("flightID")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	var body struct {
		PassengerID uint
	}

	err = c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	fmt.Println(body.PassengerID, "passssssssssssssssssssss___________")

	reservation := models.Reservation{
		ClientID:    uint(clientID),
		PassengerID: body.PassengerID,
		FlightID:    uint(flightID),
	}

	result := initializers.DB.Create(&reservation)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{"reservation": reservation})
}
func GetAllReservations(c *gin.Context) {
	clientIDStr := c.Param("id")
	clientID, err := strconv.ParseUint(clientIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	flightIDStr := c.Param("flightID")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	var reservations []models.Reservation
	result := initializers.DB.Where("client_id = ? AND flight_id = ?", clientID, flightID).Find(&reservations)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"reservations": reservations})
}

func GetReservation(c *gin.Context) {
	clientIDStr := c.Param("id")
	clientID, err := strconv.ParseUint(clientIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	flightIDStr := c.Param("flightID")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	reservationID := c.Param("ReservationID")

	var reservation models.Reservation
	result := initializers.DB.Where("client_id = ? AND flight_id = ?", clientID, flightID).First(&reservation, reservationID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	c.JSON(200, gin.H{"reservation": reservation})
}

func UpdateReservation(c *gin.Context) {
	clientIDStr := c.Param("id")
	clientID, err := strconv.ParseUint(clientIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	flightIDStr := c.Param("flightID")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	reservationID := c.Param("ReservationID")
	var body struct {
		PassengerID uint
	}
	err = c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}
	var reservation models.Reservation
	result := initializers.DB.Where("client_id = ? AND flight_id = ?", clientID, flightID).First(&reservation, reservationID)
	if result.Error != nil {
		c.Status(404)
		return
	}
	reservation.PassengerID = body.PassengerID
	result = initializers.DB.Save(&reservation)
	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{"reservation": reservation})
}

func DeleteReservation(c *gin.Context) {
	clientIDStr := c.Param("id")
	clientID, err := strconv.ParseUint(clientIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	flightIDStr := c.Param("flightID")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	reservationID := c.Param("ReservationID")

	var reservation models.Reservation
	result := initializers.DB.Where("client_id = ? AND flight_id = ?", clientID, flightID).First(&reservation, reservationID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	result = initializers.DB.Delete(&reservation)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
