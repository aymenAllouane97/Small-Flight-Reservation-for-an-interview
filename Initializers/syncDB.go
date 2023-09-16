package initializers

import models "awesomeProject/Models"

func SyncDB() {
	DB.AutoMigrate(&models.Client{})
	DB.AutoMigrate(&models.City{})
	DB.AutoMigrate(&models.Airport{})
	DB.AutoMigrate(&models.Passenger{})
	DB.AutoMigrate(&models.Flight{})
	DB.AutoMigrate(&models.Company{})
	DB.AutoMigrate(&models.Reservation{})
	DB.AutoMigrate(&models.Stopover{})
}
