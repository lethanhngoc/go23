package sorter

import (
	"sort"
	"strconv"
)

func SortInt(data []int) {
	sort.Ints(data)
}

func SortFloat(data []float64) {
	sort.Float64s(data)
}

func SortString(data []string) {
	sort.Strings(data)
}

func SortMixType(stringArr []string) ([]interface{}, error) {
	var dataString []string
	var dataFloat []float64
	var result []interface{}
	for _, data := range stringArr {
		value, error := strconv.ParseFloat(data, 64)
		if error != nil {
			dataString = append(dataString, data)
		} else {
			dataFloat = append(dataFloat, value)
		}
	}
	sort.Float64s(dataFloat)
	sort.Strings(dataString)
	for _, val := range dataFloat {
		result = append(result, val)
	}
	for _, val := range dataString {
		result = append(result, val)
	}
	return result, nil
}
