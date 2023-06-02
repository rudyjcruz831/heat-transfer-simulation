package services

import (
	"fmt"
	"time"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
)

type pumpService struct {
	pump *models.Pump
}

// Create a NewPumpService is a factory function for
// initializing a PumpService with its repository layer dependencies
func NewPumpService(pump *models.Pump) models.PumpService {
	return &pumpService{
		pump: pump,
	}
}

// Get the flow rate of pump
func (p *pumpService) GetFlowRate() float64 {
	return p.pump.FlowRate
}

func (p *pumpService) PumpWater() {
	// Pump water from the storage tank to the solar panle
	waterFlow := p.pump.FlowRate // Get the desired flow rate from the pump

	// calcualate the volume of water to be pump based on the flow rate
	volume := waterFlow * time.Second.Seconds() // Assuming time.Second is the duration for 1 second

	fmt.Printf("Pumping %.2f liters of water per seocnd...\n", volume)

	fmt.Println("Water pumped successfully")
}

func (p *pumpService) MonitorWaterFlow() {
	fmt.Println("Water flow rate : ", p.pump.FlowRate, "liters per minute")
}
