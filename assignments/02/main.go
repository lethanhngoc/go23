package main

import (
	"02/parser"
	"02/sorter"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Invalid input")
		return
	}
	sortType := args[0]
	data := args[1:]
	switch sortType {
	case "-int":
		result, err := parser.ParseInt(data)
		if err != nil {
			fmt.Printf("Can not parse and sort type %s and data %s\n", sortType, result)
		}
		sorter.SortInt(result)
		fmt.Println(result)
		break
	case "-float":
		result, err := parser.ParseFloat(data)
		if err != nil {
			fmt.Printf("Can not parse and sort type %s and data %s\n", sortType, result)
		}
		sorter.SortFloat(result)
		fmt.Println(result)
		break
	case "-string":
		sorter.SortString(data)
		fmt.Println(data)
		break
	case "-mix":
		result, err := sorter.SortMixType(data)
		if err != nil {
			fmt.Printf("Can not parse and sort type %s and data %s\n", sortType, result)
		}
		fmt.Println(result)
		break
	default:
		fmt.Println("Sort type must be [-int, -float, -string, -mix]")
	}
}
