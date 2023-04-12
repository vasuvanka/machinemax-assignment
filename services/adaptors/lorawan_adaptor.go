package adaptors

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/vasuvanka/machinemax-assignment/configuration"
	"github.com/vasuvanka/machinemax-assignment/definitions"
	"github.com/vasuvanka/machinemax-assignment/services/external"
)

type LoRaWANAdaptor struct {
	Client *external.HttpClient
}

func NewLoRaWANAdaptor() *LoRaWANAdaptor {
	return &LoRaWANAdaptor{
		Client: external.NewHttpClient(),
	}
}

func (adaptor *LoRaWANAdaptor) Register(
	ctx context.Context,
	requestDTO definitions.DevEUIRegistrationRequestDTO,
) (*definitions.DevEUIRegistrationResponseDTO, error) {
	conf := configuration.GetConfig()
	url := fmt.Sprintf("%s/sensor-onboarding-sample", conf.LoRaWANHost)

	httpHeader := http.Header{}
	httpHeader.Add("Content-Type", "application/json")

	resp, err := adaptor.Client.Post(url, requestDTO, httpHeader)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseDTO := definitions.DevEUIRegistrationResponseDTO{
		DevEUI:  requestDTO.DevEUI,
		State:   definitions.RegistrationStateFailure,
		Message: string(respBody),
	}

	switch resp.StatusCode {
	case http.StatusUnprocessableEntity:
		responseDTO.State = definitions.RegistrationStateDuplicate
	case http.StatusOK:
		responseDTO.State = definitions.RegistrationStateSuccess
	}
	return &responseDTO, nil
}
