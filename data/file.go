package data

import (
	"bufio"
	"os"
	"strconv"
)

func ReadDataFromFile(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	estimatedLines := fileInfo.Size() / 20 // Estimate for line count
	data := make([]float64, 0, estimatedLines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	return data, scanner.Err()
}
