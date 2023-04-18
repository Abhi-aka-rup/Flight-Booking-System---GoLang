package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flight-system/src/controllers"
	"flight-system/src/databaseInit"
	"flight-system/src/repository"
	"flight-system/src/services"
)

var db *gorm.DB

func main() {
	db := databaseInit.InitDb()

	flightRepo := repository.NewFlightRepository(db)
	flightService := services.NewFlightService(flightRepo)
	flightCtrl := controllers.NewFlightController(flightService)

	router := gin.Default()
	router.GET("/api/flights", flightCtrl.GetAllFlights)
	router.GET("/api/flights/:id", flightCtrl.GetFlightById)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
