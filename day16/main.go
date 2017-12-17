package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func dance(move string, dancers []string) []string {
	args := move[1:]
	switch move[0] {
	case 's':
		num, _ := strconv.Atoi(args)
		split := len(dancers) - num
		dancers = append(dancers[split:], dancers[:split]...)
	case 'x':
		t := strings.Split(args, "/")
		A, _ := strconv.Atoi(t[0])
		B, _ := strconv.Atoi(t[1])
		dancers[A], dancers[B] = dancers[B], dancers[A]
	case 'p':
		t := strings.Split(args, "/")
		A := indexOf(dancers, t[0])
		B := indexOf(dancers, t[1])
		dancers[A], dancers[B] = dancers[B], dancers[A]
	default:
		panic("Move not recognized")
	}
	//fmt.Println(move)
	//fmt.Println(dancers)
	return dancers
}

func sliceToString(slice []string) string {
	return strings.Join(slice, "")
}

func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func makeDance(dancers []string, moves []string, iterations int) string {
	solutions := []string{}
	for i := 0; i <= iterations; i++ {
		if indexOf(solutions, sliceToString(dancers)) == -1 {
			solutions = append(solutions, sliceToString(dancers))
			for _, move := range moves {
				dancers = dance(move, dancers)
			}
		} else {
			return solutions[iterations%i]
		}
	}
	return solutions[1]
}

func main() {
	dancers := []string{}
	for i := 0; i < 16; i++ {
		letter := string(97 + i)
		dancers = append(dancers, letter)
	}
	inputFile, _ := ioutil.ReadFile("input.txt")

	moves := strings.Split(string(inputFile), ",")
	sol1 := makeDance(append([]string{}, dancers...), moves, 1)
	fmt.Println("Solution 1 is:", sol1)
	sol2 := makeDance(append([]string{}, dancers...), moves, 1000000000)
	fmt.Println("Solution 2 is:", sol2)

}
