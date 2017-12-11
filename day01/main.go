package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func solveCaptcha(puzzleInput []int, step int) int {
	sum := 0
	for index, value := range puzzleInput {
		nextIndex := (index + step) % len(puzzleInput)
		if value == puzzleInput[nextIndex] {
			sum += value
		}
	}
	return sum
}

func main() {
	inputFile, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	stringInput := strings.Split(string(inputFile), "")
	puzzleInput := make([]int, len(stringInput))
	for index, value := range stringInput {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		puzzleInput[index] = i
	}

	// Phase 1
	fmt.Println(solveCaptcha(puzzleInput,1))

	// Phase 2
	fmt.Println(solveCaptcha(puzzleInput, len(puzzleInput) / 2))
}