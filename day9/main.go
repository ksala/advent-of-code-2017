package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func tokenize(line string) []string {
	return strings.Split(line, "")
}

func removeBangs(tokens []string) []string {
	cleanTokens := []string{}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "!":
			i++
		default:
			cleanTokens = append(cleanTokens, tokens[i])
		}
	}
	return cleanTokens
}

func removeGarbage(tokens []string) ([]string, int) {
	cleanTokens := []string{}
	inGarbage := false
	garbageRemoved := 0
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == ">" {
			inGarbage = false
		}
		if inGarbage {
			garbageRemoved = garbageRemoved + 1
		}
		if tokens[i] == "<" {
			inGarbage = true
		}
		if inGarbage == false {
			cleanTokens = append(cleanTokens, tokens[i])
		}
	}
	return cleanTokens, garbageRemoved
}

func calcScore(tokens []string) int {
	score := 0
	value := 1
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "{":
			score = score + value
			value = value + 1
		case "}":
			value = value - 1
		}
	}
	return score
}

func main() {
	inputFile, _ := ioutil.ReadFile("input.txt")
	tokens := tokenize(string(inputFile))
	cleanTokens, garbageRemoved := removeGarbage(removeBangs(tokens))
	fmt.Printf("Total score: %d\n", calcScore(cleanTokens))
	fmt.Printf("Garbage removed: %d\n", garbageRemoved)
}
