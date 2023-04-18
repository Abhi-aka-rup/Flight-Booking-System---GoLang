package services

import (
	"errors"
	"flight-system/src/entities"
	"flight-system/src/interfaces"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type FlightService struct {
	FlightRepo interfaces.IFlightRepository
}

func NewFlightService(flightRepo interfaces.IFlightRepository) interfaces.IFlightService {
	return FlightService{FlightRepo: flightRepo}
}

func (service FlightService) GetAllFlights() ([]entities.Flight, int, error) {
	flights, err := service.FlightRepo.GetAllFlights()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.Flight{}, http.StatusNotFound, errors.New("No flight record exist")
		} else {
			return []entities.Flight{}, http.StatusInternalServerError, err
		}
	}

	return flights, http.StatusOK, nil
}

func (service FlightService) GetFlightById(idStr string) (*entities.Flight, int, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	flight, err := service.FlightRepo.GetFlightById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("Flight not found")
		} else {
			return nil, http.StatusInternalServerError, err
		}
	}

	return flight, http.StatusOK, nil
}
