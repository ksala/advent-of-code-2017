package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func minMaxSlice(slice []int) (int, int) {
	min := slice[0]
	max := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	inputFile, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	stringInput := strings.Split(strings.TrimSuffix(string(inputFile), "\n"), "\n")

	spreadsheet := make([][]int, len(stringInput))
	for rowIndex, rowValue := range stringInput {
		row := strings.Split(rowValue, "\t")
		spreadsheet[rowIndex] = make([]int, len(row))
		for cellIndex, cellValue := range row {
			i, err := strconv.Atoi(cellValue)
			if err != nil {
				panic(err)
			}
			spreadsheet[rowIndex][cellIndex] = i
		}
	}

	checkSumPart1 := 0
	checkSumPart2 := 0
	for _, rowValue := range spreadsheet {
		min, max := minMaxSlice(rowValue)
		checkSumPart1 += max - min
		for _, cellValue1 := range rowValue {
			for _, cellValue2 := range rowValue {
				if (cellValue1 != cellValue2) && (cellValue1%cellValue2 == 0) {
					checkSumPart2 += cellValue1 / cellValue2
				}
			}
		}
	}

	fmt.Println(checkSumPart1)
	fmt.Println(checkSumPart2)
}
