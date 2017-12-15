package main

import (
	"fmt"
	"strconv"
)

const (
	// Generator A
	factorA = 16807
	//startValueA = 65
	startValueA = 699
	// Generator B
	factorB = 48271
	//startValueB = 8921
	startValueB = 124
)

func generate(previousValue int, factor int) int {
	return (previousValue * factor) % 2147483647
}

func intToBin(value int) string {
	s := strconv.FormatInt(int64(value), 2)
	s = fmt.Sprintf("%032s", s)
	return s
}

func generator(startValue int, factor int, multiple int, c chan int, done chan struct{}) {
	value := startValue
	defer close(c)
	for {
		value = generate(value, factor)
		if (value % multiple) == 0 {
			select {
			case c <- value:
				// pass
			case <-done:
				return
			}
		}
	}
}

func main() {
	// Part 1
	var matches1 int
	cA := make(chan int, 64)
	cB := make(chan int, 64)
	done1 := make(chan struct{})
	go generator(startValueA, factorA, 1, cA, done1)
	go generator(startValueB, factorB, 1, cB, done1)
	for i := 0; i <= 40000000; i++ {
		if intToBin(<-cA)[16:32] == intToBin(<-cB)[16:32] {
			matches1++
		}
	}
	close(done1)
	fmt.Printf("Matches first part: %d\n", matches1)

	// Part 2
	var matches2 int
	c2A := make(chan int, 64)
	c2B := make(chan int, 64)
	done2 := make(chan struct{})
	go generator(startValueA, factorA, 4, c2A, done2)
	go generator(startValueB, factorB, 8, c2B, done2)
	for i := 0; i <= 5000000; i++ {
		if intToBin(<-c2A)[16:32] == intToBin(<-c2B)[16:32] {
			matches2++
		}
	}
	close(done2)
	fmt.Printf("Matches second part: %d\n", matches2)
}
