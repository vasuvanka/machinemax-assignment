package utils

import (
	"testing"
)

func TestGenerateDevEUIBatch(t *testing.T) {
	// Test case 1
	limit := 10
	result, err := GenerateDevEUIBatch(limit)
	if err != nil {
		t.Errorf("Error while generating DevEUI batch: %v", err)
	}
	if len(result) != limit {
		t.Errorf("Expected length of result to be %d but got %d", limit, len(result))
	}
}
