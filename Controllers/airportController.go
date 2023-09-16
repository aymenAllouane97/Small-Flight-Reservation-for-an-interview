package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateAirport(c *gin.Context) {
	var body struct {
		Name   string `json:"name" binding:"required"`
		CityID uint   `json:"city_id" binding:"required"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var city models.City
	cityResult := initializers.DB.First(&city, body.CityID)
	if cityResult.Error != nil {
		c.JSON(404, gin.H{"error": "City not found"})
		return
	}

	airport := models.Airport{Name: body.Name, CityID: body.CityID}

	result := initializers.DB.Create(&airport)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{"airport": airport})
}

func GetAllAirports(c *gin.Context) {
	var airports []models.Airport

	result := initializers.DB.Find(&airports)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"airports": airports})
}

func GetAirport(c *gin.Context) {
	airportID := c.Param("id")

	var airport models.Airport
	result := initializers.DB.First(&airport, airportID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	c.JSON(200, gin.H{"airport": airport})
}

func GetPassengersOdAirportByDate(c *gin.Context) {
	airportID := c.Param("id")
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	var passengers []string

	result := initializers.DB.Raw("SELECT passengers.name FROM passengers JOIN reservations ON passengers.id = reservations.passenger_id JOIN flights ON reservations.flight_id = flights.id WHERE flights.departure_date = ? AND flights.departure_airport_id = ?", date, airportID).Pluck("name", &passengers)

	if result.Error != nil {
		c.Status(404)
		return
	}
	fmt.Println(result, "lisssssssssssssssst")

	c.JSON(200, gin.H{"passengers": passengers})
}
func GetCompaniesOnAirportByDate(c *gin.Context) {
	airportID := c.Param("id")
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	fmt.Println(date, " the dateee@@@@@@@@@@@@@@@@@@@2")
	var companies []string
	result := initializers.DB.Raw("SELECT distinct companies.name FROM companies JOIN flights ON companies.id = flights.company_id  WHERE( flights.departure_date = ? or flights.arrival_date=? )AND flights.departure_airport_id = ?  ", date, date, airportID).Pluck("name", &companies)

	if result.Error != nil {
		c.Status(404)
		return
	}
	fmt.Println(result, "lisssssssssssssssst")

	c.JSON(200, gin.H{"companies": companies})
}

func UpdateAirport(c *gin.Context) {
	airportID := c.Param("id")

	var body struct {
		Name   string `json:"name"`
		CityID uint   `json:"city_id"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the associated city exists
	var city models.City
	cityResult := initializers.DB.First(&city, body.CityID)
	if cityResult.Error != nil {
		c.JSON(404, gin.H{"error": "City not found"})
		return
	}

	var airport models.Airport
	result := initializers.DB.First(&airport, airportID)
	if result.Error != nil {
		c.Status(404)
		return
	}

	airport.Name = body.Name
	airport.CityID = body.CityID
	result = initializers.DB.Save(&airport)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"airport": airport})
}

func DeleteAirport(c *gin.Context) {
	airportID := c.Param("id")

	var airport models.Airport
	result := initializers.DB.First(&airport, airportID)

	if result.Error != nil {
		c.Status(404) // Not Found
		return
	}

	result = initializers.DB.Delete(&airport)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
