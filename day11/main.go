package main

import (
	"strings"
	"fmt"
	"io/ioutil"
)

type hex struct {
	x int
	y int
	z int
}

func intAbs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func findMaxInt(series []int) int {
	max := series[0]
	for _, value := range series {
		if value > max {
			max = value
		}
	}
	return max
}

func moveHex(direction string, h hex) hex {
	switch direction {
	case "ne":
		h.x++
		h.z--
	case "n":
		h.y++
		h.z--
	case "nw":
		h.x--
		h.y++
	case "sw":
		h.x--
		h.z++
	case "s":
		h.y--
		h.z++
	case "se":
		h.x++
		h.y--
	}
	return h
}

func makeHex(x, y, z int) hex {
	h := hex{x: x, y: y, z: z}
	return h
}

func manhattanDistance(h1 hex, h2 hex) int {
	return (intAbs(h1.x-h2.x) + intAbs(h1.y-h2.y) + intAbs(h1.z-h2.z)) / 2
}

func main() {
	inputFile, _ := ioutil.ReadFile("input.txt")
	directions := strings.Split(string(inputFile), ",")
	//directions := []string{"ne", "ne", "s", "s"}
	start := makeHex(0, 0, 0)
	hex := makeHex(0, 0, 0)
	distances := []int{}
	for _, direction := range directions {
		hex = moveHex(direction, hex)
		distances = append(distances, manhattanDistance(hex, start))
	}
	fmt.Printf("Distance from start: %d\n", manhattanDistance(hex, start))
	fmt.Printf("Maximum distance traveled: %d\n", findMaxInt(distances))
}
