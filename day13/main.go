package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type layer struct {
	depth           int
	scannerPosition int
	direction       int
}

type firewalls []*layer

func (l *layer) moveScanner() {
	// Boundary check
	if l.scannerPosition >= (l.depth-1) || l.scannerPosition == 0 {
		l.direction = -l.direction // Invert direction of the scanner
	}
	l.scannerPosition += l.direction
}

func newLayer(depth int) *layer {
	layer := &layer{depth: depth, direction: -1}
	return layer
}

func resetFirewall(f firewalls) {
	for _, l := range f {
		l.scannerPosition = 0
		l.direction = -1
	}
}

func findMax(m map[int]int) int {
	max := 0
	for k := range m {
		if k > max {
			max = k
		}
	}
	return max
}

func printState(fs firewalls) {
	fmt.Println("==========")
	for _, f := range fs {
		for i := 0; i < f.depth; i++ {
			if f.scannerPosition == i {
				fmt.Printf("[S] ")
			} else {
				fmt.Printf("[ ] ")
			}
		}
		fmt.Println()
	}
	fmt.Println("==========")
}

func main() {
	//depths := map[int]int{0: 3, 1: 2, 4: 4, 6: 4}

	depths := map[int]int{}
	inputFile, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(inputFile), "\n") {
		lineSplit := strings.Split(line, ":")
		num, _ := strconv.Atoi(lineSplit[0])
		depth, _ := strconv.Atoi(strings.TrimSpace(lineSplit[1]))
		depths[num] = depth
	}

	var firewall firewalls
	for i := 0; i <= findMax(depths); i++ {
		firewall = append(firewall, newLayer(depths[i]))
	}

	severity := -1
	cur := 0
	delay := -1
	caught := true
	for caught {
		// reset
		resetFirewall(firewall)
		//printState(firewall)
		cur = -1
		severity = 0
		caught = false
		// Calc delay
		delay++
		for i := 0; i < delay; i++ {
			for _, f := range firewall {
				f.moveScanner()
			}
		}
		//fmt.Printf("Delay: %d\n", delay)
		//printState(firewall)
		// Run
		for i := 0; i < len(firewall); i++ {
			cur++
			if firewall[cur].scannerPosition == 0 && firewall[cur].depth != 0 {
				caught = true
				severity = severity + (firewall[cur].depth * cur)
			}
			for _, f := range firewall {
				f.moveScanner()
			}
		}
		//fmt.Println(severity)
		//delay++
	}
	fmt.Printf("Severity is %d\n", severity)
	fmt.Println(delay)
}
