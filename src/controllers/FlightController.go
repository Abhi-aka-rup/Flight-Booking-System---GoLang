package controllers

import (
	"flight-system/src/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	FlightService interfaces.IFlightService
}

func NewFlightController(flightService interfaces.IFlightService) FlightController {
	return FlightController{FlightService: flightService}
}

func (controller FlightController) GetAllFlights(context *gin.Context) {
	flights, status, err := controller.FlightService.GetAllFlights()
	if err != nil {
		context.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, flights)
}

func (controller FlightController) GetFlightById(context *gin.Context) {
	flights, status, err := controller.FlightService.GetFlightById(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, flights)
}
