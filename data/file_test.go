package data

import (
	"os"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	// Create a temporary data file for testing
	tempFile, err := os.CreateTemp("", "testdata*.txt")
	if err != nil {
		t.Fatal("Could not create temp file:", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the file
	testData := "189\n113\n121\n114\n145\n110\n"
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Fatal("Could not write to temp file:", err)
	}
	tempFile.Close()

	// Test the ReadDataFromFile function
	data, err := ReadDataFromFile(tempFile.Name())
	if err != nil {
		t.Fatal("Error reading data from file:", err)
	}

	expectedData := []float64{189, 113, 121, 114, 145, 110}
	if len(data) != len(expectedData) {
		t.Fatalf("Expected %d elements, got %d", len(expectedData), len(data))
	}

	for i, v := range expectedData {
		if data[i] != v {
			t.Errorf("Expected value at index %d to be %f, got %f", i, v, data[i])
		}
	}
}
