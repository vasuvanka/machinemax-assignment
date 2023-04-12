package external

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHttpClient struct{}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	// Create a mock response
	response := httptest.NewRecorder()
	response.WriteHeader(http.StatusOK)
	return response.Result(), nil
}

func TestPost(t *testing.T) {
	// Test case 1
	httpClient := &HttpClient{
		Client: &MockHttpClient{},
	}
	url := "https://example.com"
	body := map[string]string{
		"key": "value",
	}
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	response, err := httpClient.Post(url, body, headers)
	if err != nil {
		t.Errorf("Error while making POST request: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Error unexpected http status code from POST request: %s", response.Status)
	}
}
