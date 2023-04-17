package services

import (
	"flight-system/src/entities"
	"flight-system/src/interfaces"
)

type FlightService struct {
	FlightRepo interfaces.IFlightRepository
}

func NewFlightService(flightRepo interfaces.IFlightRepository) interfaces.IFlightService {
	return FlightService{FlightRepo: flightRepo}
}

func (service FlightService) GetAllFlights() ([]entities.Flight, error) {
	return service.FlightRepo.GetAllFlights()
}
