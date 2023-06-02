package mocks

import "github.com/stretchr/testify/mock"

// MockStorageTankService is a mock implementation of the StorageTankService interface
type MockStorageTankService struct {
	mock.Mock
}

// GetVolume mocks the GetVolume method of the StorageTankService interface
func (m *MockStorageTankService) GetVolume() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

// GetTemperature mocks the GetTemperature method of the StorageTankService interface
func (m *MockStorageTankService) GetTemperature() float64 {
	args := m.Called()
	return args.Get(0).(float64)
}

// UpdateTemperature mocks the UpdateTemperature method of the StorageTankService interface
func (m *MockStorageTankService) UpdateTemperature(temperature float64) {
	m.Called(temperature)
}
