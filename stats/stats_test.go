package stats

import (
	"math"
	"testing"
)

func TestCalculateLinearRegressionAndPearson(t *testing.T) {
	// Test data
	x := []float64{0, 1, 2, 3, 4, 5}
	y := []float64{189, 113, 121, 114, 145, 110}

	// Call the function
	beta1, beta0, r := CalculateLinearRegressionAndPearson(x, y)

	// Known expected results (pre-calculated for this data)
	expectedBeta0 := 153.857143
	expectedBeta1 := -8.742857
	expectedR := -0.5330331012

	// Allowable delta for floating-point comparison
	const delta = 0.000001

	// Assert the results
	if math.Abs(beta1-expectedBeta1) > delta {
		t.Errorf("Expected Beta1 to be %f, got %f", expectedBeta1, beta1)
	}
	if math.Abs(beta0-expectedBeta0) > delta {
		t.Errorf("Expected Beta0 to be %f, got %f", expectedBeta0, beta0)
	}
	if math.Abs(r-expectedR) > delta {
		t.Errorf("Expected Pearson Correlation Coefficient to be %f, got %f", expectedR, r)
	}
}
