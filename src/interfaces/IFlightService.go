package interfaces

import "flight-system/src/entities"

type IFlightService interface {
	GetAllFlights() ([]entities.Flight, error)
}
