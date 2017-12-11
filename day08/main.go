package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type state struct {
	registers map[string]int
	biggestValue int
}

func (s state) setRegister(register string, value int) {
	s.registers[register] = value
}

func (s state) getRegister(register string) int {
	return s.registers[register]
}

func (s state) getBiggestRegister() int {
	biggestValue := 0
	for _, value := range s.registers {
		if value > biggestValue {
			biggestValue = value
		}
	}
	return biggestValue
}

func (s *state) setBiggestValue(value int) {
	s.biggestValue = value
}

func (s *state) runInstruction(instruction []string) {
	register := instruction[0]
	op := instruction[1]
	quantity, _ := strconv.Atoi(instruction[2])
	switch op {
	case "inc":
		s.setRegister(register, s.getRegister(register)+quantity)
	case "dec":
		s.setRegister(register, s.getRegister(register)-quantity)
	}
}

func (s state) evaluateCondition(condition []string) bool {
	register := s.getRegister(condition[0])
	op := condition[1]
	quantity, _ := strconv.Atoi(condition[2])
	switch op {
	case "==":
		return register == quantity
	case ">":
		return register > quantity
	case ">=":
		return register >= quantity
	case "<":
		return register < quantity
	case "<=":
		return register <= quantity
	case "!=":
		return register != quantity
	default:
		return false
	}
}

func newState() state {
	state := state{}
	state.registers = make(map[string]int)
	return state
}

func main() {
	state := newState()
	inputFile, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(inputFile), "\n") {
		slice := strings.Split(line, " ")
		instruction := slice[0:3]
		condition := slice[4:]
		if state.evaluateCondition(condition) {
			state.runInstruction(instruction)
			if state.getBiggestRegister() > state.biggestValue {
				state.setBiggestValue(state.getBiggestRegister())
			}
		}
	}

	fmt.Printf("Biggest value in a register at end of execution: %d\n", state.getBiggestRegister())
	fmt.Printf("Biggest value ever in a register: %d\n ", state.biggestValue)
}
