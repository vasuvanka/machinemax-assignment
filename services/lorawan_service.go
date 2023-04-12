package services

import (
	"context"

	"github.com/vasuvanka/machinemax-assignment/definitions"
	"github.com/vasuvanka/machinemax-assignment/services/adaptors"
)

type LoRaWANService struct {
	LoRaWANAdaptor *adaptors.LoRaWANAdaptor
}

func NewLoRaWANService() *LoRaWANService {
	return &LoRaWANService{
		LoRaWANAdaptor: adaptors.NewLoRaWANAdaptor(),
	}
}

func (service *LoRaWANService) Register(
	ctx context.Context,
	devEUI definitions.DevEUI,
) (*definitions.DevEUIRegistrationResponseDTO, error) {
	requestDTO := definitions.DevEUIRegistrationRequestDTO{
		DevEUI: devEUI,
	}
	return service.LoRaWANAdaptor.Register(
		ctx,
		requestDTO,
	)
}
