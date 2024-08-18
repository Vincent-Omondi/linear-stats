package main

import (
	"fmt"
	"os"

	"github.com/Vincent-Omondi/linear-stats/data"
	"github.com/Vincent-Omondi/linear-stats/stats"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run linear-stats.go <data_file>")
		return
	}

	fileName := os.Args[1]
	data, err := data.ReadDataFromFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	n := len(data)
	if n == 0 {
		fmt.Println("No data to process")
		return
	}

	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
	}

	beta1, beta0, r := stats.CalculateLinearRegressionAndPearson(x, data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", beta1, beta0)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}




