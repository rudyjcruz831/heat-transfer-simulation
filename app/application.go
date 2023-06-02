package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
	"github.com/rudyjcruz831/heat-transfer-simulation/services"
)

const (
	HeatTransferCoefficient = 0.1
)

func StartApplication() {

	// this is for when some hits CTRL+C to stop continues simulation
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		fmt.Println("Simulation interrupted. Stopping...")
		os.Exit(0)
	}()

	// log.Println("Starting server... ")
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

	pump := &models.Pump{
		FlowRate: 10.0,
	}

	system := &models.System{
		SolarRadiation: 1.0, // Set an initial value for solar radiation
		// AmbientTemperature: 0.0, // Set an initial value for ambient temperature
	}

	solarPaneService := services.NewSolarPanelService(solarPanle)
	storageTankService := services.NewStorageTankService(storageTank)
	pumpService := services.NewPumpService(pump)
	systemService := services.NewSystemService(system, solarPaneService, storageTankService, pumpService)

	fmt.Println("Starting Simulation...")
	// time.Second.Seconds()
	systemService.Simulate(system.SolarRadiation, HeatTransferCoefficient, *solarPanle)
}
