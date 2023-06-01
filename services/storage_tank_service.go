package services

import "github.com/rudyjcruz831/heat-transfer-simulation/models"

// stroageTankService acts as a struct for injecting an implementation of storageTank struct
// for use in service methods
type storageTankService struct {
	storageTank *models.StorageTank
}

// Create a NewStorageTankSercivice is a factory function for
// initializing a StroagejTankService with its repository layer dependencies
func NewStorageTankService(storageTank *models.StorageTank) models.StorageTankService {
	return &storageTankService{
		storageTank: storageTank,
	}
}

// Get volume of tank
func (t *storageTankService) GetVolume() float64 {
	return t.storageTank.Volume
}

// Get temperature of storage tank
func (t *storageTankService) GetTemperature() float64 {
	return t.storageTank.Temperature
}

// Update temperature of storage tank
func (t *storageTankService) UpdateTemperature(temperature float64) {
	t.storageTank.Temperature = temperature
}
