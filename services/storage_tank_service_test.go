package services

import (
	"testing"

	"github.com/rudyjcruz831/heat-transfer-simulation/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestTransferHeat(t *testing.T) {
	// mockSolarPanelSrv := &mocks.MockSolarPanel{}
	storageTank := &models.StorageTank{
		Temperature: 25.0,
		Volume:      100,
	}
	mockStorageTankSrv := NewStorageTankService(storageTank)
	solarPanel := &models.SolarPanel{
		Temperature: 50.0,
		Area:        2.0,
	}
	// mockSolarPanelSrv := NewSolarPanelService(solarPanel)

	heatTransferCoefficient := 0.5

	heatTranfer := heatTransferCoefficient * 2.0 * (50.0 - 25.0)
	mockTemp := 25.0 + heatTranfer/(100.0*4.18)
	// Set up the expected temperature after transferring heat
	// expectedStorageTankTemperature := initialStorageTankTemperature + 100.0 // Example: heatTranfer = 100.0

	mockStorageTankSrv.TransferHeat(0.5, *solarPanel) // Example: heatTransferCoefficient = 0.5
	assert.Equal(t, storageTank.Temperature, mockTemp)
	// assert.Equal(t, expectedTemperature, solarPanelService.GetTemperature())
}
