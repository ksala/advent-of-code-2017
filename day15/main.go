package main

import (
	"fmt"
	"strconv"
)

const (
	// Generator A
	factorA     = 16807
	startValueA = 699
	// Generator B
	factorB     = 48271
	startValueB = 124
)

func generate(previousValue int, factor int) int {
	return (previousValue * factor) % 2147483647
}

func intToBin(value int) string {
	s := strconv.FormatInt(int64(value), 2)
	for len(s) != 32 {
		s = fmt.Sprintf("0%s", s)
	}
	return s
}

func main() {
	previousValueA := startValueA
	previousValueB := startValueB
	var matches int
	for i := 0; i <= 40000000; i++ {
		valueA := generate(previousValueA, factorA)
		valueB := generate(previousValueB, factorB)
		if intToBin(valueA)[16:32] == intToBin(valueB)[16:32] {
			matches++
		}
		previousValueA = valueA
		previousValueB = valueB
	}
	fmt.Println(matches)
}
