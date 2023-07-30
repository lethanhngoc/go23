package main

import "fmt"

func countRectangle(arr [][]int) int {
	result := 0
	rows := len(arr)
	columns := len(arr[0])
	processed := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		processed[i] = make([]bool, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if processed[i][j] {
				continue
			}
			if arr[i][j] == 1 {
				processed[i][j] = true
				for x := i; x < rows; x++ {
					for y := j; y < columns; y++ {
						if arr[x][y] != 1 {
							break
						}
						processed[x][y] = true
					}
					if arr[x][j] != 1 {
						break
					}
				}
				result++
			}
		}
	}
	return result
}

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	result := countRectangle(arr)
	fmt.Print(result)
}
