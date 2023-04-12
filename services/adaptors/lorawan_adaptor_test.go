package adaptors

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vasuvanka/machinemax-assignment/definitions"
	"github.com/vasuvanka/machinemax-assignment/services/external"
)

type MockHttpClient struct{}

func (m *MockHttpClient) Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	// Create a mock response
	response := httptest.NewRecorder()
	response.WriteHeader(http.StatusOK)
	return response.Result(), nil
}

func TestRegister(t *testing.T) {
	// Test case 1
	adaptor := &LoRaWANAdaptor{
		Client: external.NewHttpClient(),
	}
	requestDTO := definitions.DevEUIRegistrationRequestDTO{
		DevEUI: "1234567890abcdef",
	}
	_, err := adaptor.Register(context.Background(), requestDTO)
	if err != nil {
		t.Errorf("Error while registering DevEUI: %v", err)
	}
}
