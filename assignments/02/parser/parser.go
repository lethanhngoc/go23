package parser

import (
	"strconv"
)

func ParseInt(stringArr []string) ([]int, error) {
	result := make([]int, len(stringArr))
	for i, data := range stringArr {
		value, error := strconv.Atoi(data)
		if error != nil {
			return result, error
		}
		result[i] = value
	}
	return result, nil
}

func ParseFloat(stringArr []string) ([]float64, error) {
	result := make([]float64, len(stringArr))
	for i, data := range stringArr {
		value, error := strconv.ParseFloat(data, 64)
		if error != nil {
			return result, error
		}
		result[i] = value
	}
	return result, nil
}
