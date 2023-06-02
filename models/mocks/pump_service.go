package mocks

import "github.com/stretchr/testify/mock"

type MockPumpService struct {
	mock.Mock
}

func (m *MockPumpService) GetFlowRate() float64 {
	ret := m.Called()

	var r0 float64
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
