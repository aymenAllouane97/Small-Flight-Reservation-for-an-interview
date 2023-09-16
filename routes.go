package main

import (
	controllers "awesomeProject/Controllers"
	middleware "awesomeProject/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	clientRoutes := r.Group("/clients")
	clientRoutes.Use(middleware.AuthMiddleware, middleware.ClientAuthMiddleware)
	{
		clientRoutes.POST("/", controllers.CreateClient)
		clientRoutes.POST("/login", controllers.LoginClient)
		clientRoutes.GET("/", controllers.GetAllClients)
		clientRoutes.GET("/:id", controllers.GetClient)
		clientRoutes.PUT("/:id", controllers.UpdateClient)
		clientRoutes.DELETE("/:id", controllers.DeleteClient)

		reservationsRoutes := clientRoutes.Group("/:id/flights/:flightID/reservations")
		{
			reservationsRoutes.POST("/", controllers.CreateReservation)
			reservationsRoutes.GET("/", controllers.GetAllReservations)
			reservationsRoutes.GET("/:ReservationID", controllers.GetReservation)
			reservationsRoutes.PUT("/:ReservationID", controllers.UpdateReservation)
			reservationsRoutes.DELETE("/:ReservationID", controllers.DeleteReservation)
		}
	}

	passengerRoutes := r.Group("/passengers")
	passengerRoutes.Use(middleware.AuthMiddleware, middleware.ClientAuthMiddleware)
	{
		passengerRoutes.POST("/", controllers.CreatePassenger)
		passengerRoutes.GET("/", controllers.GetAllPassengers)
		passengerRoutes.GET("/:id", controllers.GetPassenger)
		passengerRoutes.PUT("/:id", controllers.UpdatePassenger)
		passengerRoutes.DELETE("/:id", controllers.DeletePassenger)
	}

	companyRoutes := r.Group("/companies")
	//companyRoutes.Use(middleware.AuthMiddleware, middleware.CompanyAuthMiddleware)
	{
		companyRoutes.POST("/", controllers.CreateCompany)
		companyRoutes.GET("/", controllers.GetAllCompanies)
		companyRoutes.POST("/login", controllers.LoginCompany)
		companyRoutes.GET("/:id", controllers.GetCompany)
		companyRoutes.PUT("/:id", controllers.UpdateCompany)
		companyRoutes.DELETE("/:id", controllers.DeleteCompany)
	}

	flightRoutes := r.Group("/flights")
	flightRoutes.Use(middleware.AuthMiddleware, middleware.CompanyAuthMiddleware)
	{
		flightRoutes.POST("/", controllers.CreateFlight)
		flightRoutes.GET("/", controllers.GetAllFlights)
		flightRoutes.GET("/:id", controllers.GetFlight)
		flightRoutes.PUT("/:id", controllers.UpdateFlight)
		flightRoutes.DELETE("/:id", controllers.DeleteFlight)
		stopoversRoute := flightRoutes.Group("/:id/stopovers")
		{
			stopoversRoute.POST("/", controllers.CreateStopover)
			stopoversRoute.GET("/:stopID", controllers.GetStopover)
			stopoversRoute.GET("/", controllers.GetAllStopovers)
			stopoversRoute.PUT("/:stopID", controllers.UpdateStopover)
			stopoversRoute.DELETE("/:stopID", controllers.DeleteStopover)
		}

	}

	cityRoutes := r.Group("/cities")
	{
		cityRoutes.POST("/", controllers.CreateCity)
		cityRoutes.GET("/", controllers.GetAllCities)
		cityRoutes.GET("/:id", controllers.GetCity)
		cityRoutes.PUT("/:id", controllers.UpdateCity)
		cityRoutes.DELETE("/:id", controllers.DeleteCity)
	}
	airportRoutes := r.Group("/airports")
	{
		airportRoutes.POST("/", controllers.CreateAirport)
		airportRoutes.GET("/", controllers.GetAllAirports)
		airportRoutes.GET("/:id", controllers.GetAirport)
		airportRoutes.GET("/:id/passengers/:date", controllers.GetPassengersOdAirportByDate)
		airportRoutes.GET("/:id/companies/:date", controllers.GetCompaniesOnAirportByDate)
		airportRoutes.PUT("/:id", controllers.UpdateAirport)
		airportRoutes.DELETE("/:id", controllers.DeleteAirport)

	}
}
