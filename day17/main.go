package main

import "fmt"

const (
	steps = 370
)

func insert(before []int, value int, after []int) []int {
	result := []int{}
	for _, v := range before {
		result = append(result, v)
	}
	result = append(result, value)
	for _, v := range after {
		result = append(result, v)
	}
	return result
}

func part1() {
	buffer := []int{0}
	var cur int
	for i := 1; i < 2018; i++ {
		cur = ((cur + steps) % i) + 1
		buffer = insert(buffer[:cur], i, buffer[cur:])
	}
	fmt.Println("Solution part 1:", buffer[cur+1])
}

func part2() {
	var cur int
	var solution int
	for i := 1; i < 50000000; i++ {
		cur = ((cur + steps) % i) + 1
		if cur == 1 {
			solution = i
		}
	}
	fmt.Println("Solution part 2:", solution)
}

func main() {
	part1()
	part2()
}
