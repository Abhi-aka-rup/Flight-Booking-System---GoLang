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

	return flights, result.Error
}

func (repo FlightRepository) GetFlightById(id uint) (*entities.Flight, error) {
	var flight entities.Flight
	result := repo.DB.First(&flight, id)

	return &flight, result.Error
}
