package interfaces

import "flight-system/src/entities"

type IFlightService interface {
	GetAllFlights() ([]entities.Flight, int, error)
	GetFlightById(idStr string) (*entities.Flight, int, error)
}
