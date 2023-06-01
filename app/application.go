package app

import (
	"log"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
	"github.com/rudyjcruz831/heat-transfer-simulation/services"
)

const (
	HeatTransferCoefficient = 0.1
)

func StartApplication() {
	log.Println("Starting server... ")
	solarPanle := &models.SolarPanel{
		Area:           2.0,
		Efficiency:     0.8,
		Temperature:    0.0,
		Degradation:    1.0,
		DustAccumulate: 0.0,
	}

	storageTank := &models.StorageTank{
		Volume:      100.00,
		Temperature: 25.0,
	}

	// pump := &models.Pump{
	// 	FlowRate: 10.0,
	// }

	system := &models.System{
		SolarRadiation: 1.0, // Set an initial value for solar radiation
		// AmbientTemperature: 0.0, // Set an initial value for ambient temperature
	}

	solarPaneService := services.NewSolarPanelService(solarPanle)
	storageTankService := services.NewStorageTankService(storageTank)
	// pumpService := services.NewPumpService(pump)
	systemService := services.NewSystemService(system, solarPaneService, storageTankService)

	systemService.Simulate(system.SolarRadiation, HeatTransferCoefficient)
}
