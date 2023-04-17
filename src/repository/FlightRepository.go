package repository

import (
	"flight-system/src/entities"
	"flight-system/src/interfaces"

	"gorm.io/gorm"
)

type FlightRepository struct {
	DB *gorm.DB
}

func NewFlightRepository(db *gorm.DB) interfaces.IFlightRepository {
	return FlightRepository{DB: db}
}

func (repo FlightRepository) GetAllFlights() ([]entities.Flight, error) {
	var flights []entities.Flight
	result := repo.DB.Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}
	return flights, nil
}
