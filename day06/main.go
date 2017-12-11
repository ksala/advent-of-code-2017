package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strSlice := strings.Split(strings.TrimSpace(string(inputFile)), "\t")
	banks := make([]int, len(strSlice))
	for index, value := range strSlice {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		banks[index] = i
	}

	previousStates := []string(nil)
	steps := 0
	previousStateIndex := 0
Loop:
	for {
		// Find bank with most blocks
		biggestBankValue := 0
		biggestBankIndex := 0
		for index, value := range banks {
			if value > biggestBankValue {
				biggestBankIndex = index
				biggestBankValue = value
			}
		}
		// Set bank to 0
		blocks := banks[biggestBankIndex]
		banks[biggestBankIndex] = 0
		// Add 1 block to each bank sequential
		next := biggestBankIndex + 1
		for x := blocks; x > 0; x-- {
			if next >= len(banks) {
				next = 0
			}
			banks[next]++
			next++
		}
		// Check if banks state is same as before
		steps++
		state := fmt.Sprint(banks)
		for index, value := range previousStates {
			if state == value {
				previousStateIndex = index
				break Loop
			}
		}
		// Add banks state to previous state
		previousStates = append(previousStates, fmt.Sprint(banks))
	}
	fmt.Println(steps)
	fmt.Println(steps - (previousStateIndex + 1))
}
