package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockSolarPanel struct {
	mock.Mock
}

func (m *MockSolarPanel) GetArea() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSolarPanel) GetEfficiency() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSolarPanel) GetTemperature() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSolarPanel) GetDegradation() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

func (m *MockSolarPanel) CaptureSolarRadiation() {
	m.Called()
}

func (m *MockSolarPanel) CaptureDustAccumulation() {
	m.Called()
}

func (m *MockSolarPanel) CaptureSolarEnergy(solarRediation float64, heatTransferCoefficient float64) {
	m.Called(solarRediation, heatTransferCoefficient)
}

func (m *MockSolarPanel) UpdateDegradation() {
	m.Called()
}
