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
	flights, err := controller.FlightService.GetAllFlights()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, flights)
}
