package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockSystemService struct {
	mock.Mock
}

func (m *MockSystemService) CaptureSolarRadiation() {
	m.Called()
}

func (m *MockSystemService) CaptureDustAccumulation() {
	m.Called()
}

func (m *MockSystemService) CaptureSolarEnergy(solarRadiation float64, heatTransferCoefficient float64) {
	m.Called(solarRadiation, heatTransferCoefficient)
}

func (m *MockSystemService) UpdateDegradation() {
	m.Called()
}

func (m *MockSystemService) GetTemperature() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSystemService) GetEfficiency() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSystemService) TransferHeat(heatTransferCoefficient float64) {
	m.Called(heatTransferCoefficient)
}

func (m *MockSystemService) MonitorWaterFlow() {
	m.Called()
}

func (m *MockSystemService) DisplayTemperatures() {
	m.Called()
}

func (m *MockSystemService) Simulate(solarRadiation float64, heatTransferCoefficient float64) {
	m.Called(solarRadiation, heatTransferCoefficient)
}
