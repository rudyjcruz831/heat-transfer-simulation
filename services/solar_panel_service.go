package services

import (
	"math/rand"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
)

// solarPanelService acts as a struct for injecting an implementation of SolarPanel struct
// for use in service methods
type solarPanelService struct {
	solarPanel *models.SolarPanel
}

// Create a NewSolarPanelService is a factory function for
// initializing a SolarPanelService with its repository layer dependencies
func NewSolarPanelService(solarPanel *models.SolarPanel) models.SolarPanelService {
	return &solarPanelService{
		solarPanel: solarPanel,
	}
}

// Get Area for SolarPanel
func (p *solarPanelService) GetArea() float64 {
	return p.solarPanel.Area
}

// Get Efficiency for Solar Panel
func (p *solarPanelService) GetEfficiency() float64 {
	return p.solarPanel.Efficiency
}

// Get Temperatur of Solar Panel
func (p *solarPanelService) GetTemperature() float64 {
	return p.solarPanel.Temperature
}

// Get Degradation for Solar Panel
func (p *solarPanelService) GetDegradation() float64 {
	return p.solarPanel.Degradation
}

// Get Dust accumulation for solar Panel
func (p *solarPanelService) GetDustAccumulation() float64 {
	return p.solarPanel.DustAccumulate
}

func (p *solarPanelService) CaptureSolarRadiation() {
	// simulate dynamic solar radiation (esample: random values between 0.5 and 1.5)
	p.solarPanel.Efficiency = p.solarPanel.Efficiency * (0.5 + rand.Float64())
}

// In this example, rand.Float64() generates a random value
// between 0 and 1, which is then multiplied by 0.1 to
// simulate dust accumulation. You can adjust the
// multiplication factor or the range of random values
// according to your specific requirements.
func (p *solarPanelService) CaptureDustAccumulation() {
	// Simulate dust accumulation on solar panel (example: random values between 0 and 0.1)
	p.solarPanel.DustAccumulate = rand.Float64() * 0.1
}

// CaptureSolarEnergy captures solar energy and increases the solar panel's temperature.
// It calculates the incident solar radiation based on the efficiency, area, and given solar radiation.
// Then, it updates the temperature using the heat transfer coefficient.
func (p *solarPanelService) CaptureSolarEnergy(solarRediation float64, heatTransferCoefficient float64) {
	incidentSolarRadiation := p.GetEfficiency() * p.GetArea() * solarRediation
	p.solarPanel.Temperature += incidentSolarRadiation / (p.GetArea() * heatTransferCoefficient)

}

func (p *solarPanelService) UpdateDegradation() {
	// simulate solar panel degradation over time
	// (example: degradation factor decreses by 0.001 per second)
	p.solarPanel.Degradation -= 0.001
}
