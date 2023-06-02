package services

import (
	"fmt"
	"time"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
)

// systemService acts as a struct for injecting an implementation of
// SolarPanel, SolarPanelService, StorageTankService struct for use in service methods
type systemService struct {
	system         *models.System
	solarPanelSrv  models.SolarPanelService
	storageTankSrv models.StorageTankService
	pumpSrv        models.PumpService
}

// Create a NewSystemService is a factory function for
// initializing a SystemService with its repository layer dependencies
func NewSystemService(system *models.System,
	solarasolarPanelService models.SolarPanelService,
	storageTankService models.StorageTankService,
	pumpService models.PumpService) models.SystemService {
	return &systemService{
		system:         system,
		solarPanelSrv:  solarasolarPanelService,
		storageTankSrv: storageTankService,
		pumpSrv:        pumpService,
	}
}

// CaptureSolarRadiation captures the solar radiation
// by invoking the corresponding method in the solar panel service.
func (s *systemService) CaptureSolarRadiation() {
	s.solarPanelSrv.CaptureSolarRadiation()
}

// CaptureDustAccumulation captures the dust accumulation on the solar
// panel by invoking the corresponding method in the solar panel service.
func (s *systemService) CaptureDustAccumulation() {
	s.solarPanelSrv.CaptureDustAccumulation()
}

// CaptureSolarEnergy captures the solar energy based on the provided
// solar radiation and heat transfer coefficient.
// It invokes the corresponding method in the solar panel service.
func (s *systemService) CaptureSolarEnergy(solarRadiation float64, heatTransferCoefficient float64) {
	s.solarPanelSrv.CaptureSolarEnergy(solarRadiation, heatTransferCoefficient)
}

// UpdateDegradation updates the degradation of the solar panel by invoking
// the corresponding method in the solar panel service.
func (s *systemService) UpdateDegradation() {
	s.solarPanelSrv.UpdateDegradation()
}

// GetTemperature retrieves the temperature of the storage tank
func (s *systemService) GetTemperature() float64 {
	return s.storageTankSrv.GetTemperature()
}

// GetEfficiency retrieves the efficiency of the solar panel
func (s *systemService) GetEfficiency() float64 {
	return s.solarPanelSrv.GetEfficiency()
}

// Transfer heat from SolarPanel to StorageTank
func (s *systemService) TransferHeat(heatTransferCoefficient float64, solarPanel models.SolarPanel) {
	s.storageTankSrv.TransferHeat(heatTransferCoefficient, solarPanel)
}

func (s *systemService) MonitorWaterFlow() {
	s.pumpSrv.MonitorWaterFlow()
}

// DisplayTemperatures prints the temperatures and other relevant data to the console
func (s *systemService) DisplayTemperatures() {
	fmt.Printf("Solar Panel Temperature: %.2f°C\n", s.solarPanelSrv.GetTemperature())
	fmt.Printf("Storage Tank Temperature: %.2f°C\n", s.storageTankSrv.GetTemperature())
	fmt.Printf("Solar Radiation: %.2f\n", s.system.SolarRadiation)
	fmt.Printf("Solar Panel Efficiency: %.2f\n", s.GetEfficiency())
	fmt.Printf("Solar Panel Degradation: %.3f\n", s.solarPanelSrv.GetDegradation())
	fmt.Println("----------------------------------------")
}

func (s *systemService) Simulate(solarRadiation float64,
	heatTransferCoefficient float64,
	solarPanel models.SolarPanel) {
	for {
		s.CaptureSolarRadiation()
		// s.CaptureAmbientTemperature()
		s.CaptureDustAccumulation()
		s.CaptureSolarEnergy(solarRadiation, heatTransferCoefficient)
		s.TransferHeat(heatTransferCoefficient, solarPanel)
		s.UpdateDegradation()
		s.MonitorWaterFlow()
		s.DisplayTemperatures()

		time.Sleep(time.Second)
	}
}
