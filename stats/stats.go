package stats

import "math"

func CalculateLinearRegressionAndPearson(x, y []float64) (float64, float64, float64) {
	n := float64(len(x))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
		sumY2 += y[i] * y[i]
	}

	beta1 := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	beta0 := (sumY - beta1*sumX) / n

	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	r := 0.0
	if denominator != 0 {
		r = numerator / denominator
	}

	return beta1, beta0, r
}
