package models

// SystemService interface represetns a system simulation of
// Solar System , Pump, and Storage Tank
type SystemService interface {
	CaptureSolarRadiation()
	CaptureDustAccumulation()
	CaptureSolarEnergy(float64, float64)
	UpdateDegradation()
	GetTemperature() float64
	GetEfficiency() float64
	DisplayTemperatures()
	TransferHeat(float64, SolarPanel)
	Simulate(float64, float64, SolarPanel)
}

// SolarPanelSevice interface represents a stroage tank service
type SolarPanelService interface {
	GetArea() float64
	GetEfficiency() float64
	GetTemperature() float64
	GetDegradation() float64
	GetDustAccumulation() float64
	CaptureSolarRadiation()
	CaptureDustAccumulation()
	CaptureSolarEnergy(solarRadiation float64,
		heatTransferCoefficient float64)
	UpdateDegradation()
}

// StorageTankService interface represents a stroage tank service
type StorageTankService interface {
	GetTemperature() float64
	GetVolume() float64
	UpdateTemperature(float64)
	TransferHeat(float64, SolarPanel)
}

// PumpService interface represents a pump component
type PumpService interface {
	GetFlowRate() float64
	PumpWater()
	MonitorWaterFlow()
}
