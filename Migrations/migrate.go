package main

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func SyncDB() {
	initializers.DB.AutoMigrate(&models.Client{})
	initializers.DB.AutoMigrate(&models.City{})
	initializers.DB.AutoMigrate(&models.Airport{})
	initializers.DB.AutoMigrate(&models.Passenger{})
	initializers.DB.AutoMigrate(&models.Flight{})
	initializers.DB.AutoMigrate(&models.Company{})
	initializers.DB.AutoMigrate(&models.Reservation{})
	initializers.DB.AutoMigrate(&models.Stopover{})
}
