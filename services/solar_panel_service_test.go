package services

import (
	"testing"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
	"github.com/rudyjcruz831/heat-transfer-simulation/models/mocks"
	"gopkg.in/go-playground/assert.v1"
)

func TestCaptureSolarEnergy(t *testing.T) {

	// Set the initial solarPanel
	solarPanel := &models.SolarPanel{
		Area:        2.0,
		Efficiency:  0.8,
		Temperature: 25.0,
	}
	solarRadiation := 1000.0
	heatTransferCoefficient := 0.5
	solarPanelService := NewSolarPanelService(solarPanel)

	solarPanelService.CaptureSolarEnergy(solarRadiation, heatTransferCoefficient)

	// Check the updated temperature
	incidentSolarRadiation := solarPanel.Efficiency * solarPanel.Area * solarRadiation
	expectedTemperature := 25.0 + incidentSolarRadiation/(2.0*0.5)
	assert.Equal(t, expectedTemperature, solarPanelService.GetTemperature())

}

func TestUpdateDegradation(t *testing.T) {
	// Create a mock solar panel
	mockSolarPanel := &mocks.MockSolarPanel{}

	solarPanel := &models.SolarPanel{
		Degradation: 0.5,
	}
	solarPanelService := NewSolarPanelService(solarPanel)
	solarPanelService.UpdateDegradation()

	// Set the expected values for the UpdateDegradation method
	mockSolarPanel.On("UpdateDegradation").Return()

	// Update degradation
	mockSolarPanel.UpdateDegradation()

	// Check the updated degradation
	expectedDegradation := 0.5 - 0.001
	assert.Equal(t, expectedDegradation, solarPanel.Degradation)

	// Assert that the expected methods were called
	mockSolarPanel.AssertExpectations(t)
}
