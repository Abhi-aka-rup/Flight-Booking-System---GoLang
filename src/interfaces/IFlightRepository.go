package interfaces

import "flight-system/src/entities"

type IFlightRepository interface {
	GetAllFlights() ([]entities.Flight, error)
	GetFlightById(id uint) (*entities.Flight, error)
}
