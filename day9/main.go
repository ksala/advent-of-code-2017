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
	for i := 0; i < len(tokens); i ++ {
		if tokens[i] == "<" {
			inGarbage = true
		}
		if inGarbage == false {
			cleanTokens = append(cleanTokens, tokens[i])
		}
		if tokens[i] == ">" && inGarbage {
			inGarbage = false
		}
	}
	return cleanTokens, garbageRemoved
}

func calcScore(tokens []string) int {
	score := 0
	value := 1
	for i := 0; i < len(tokens); i ++ {
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

func popSlice(slice []string) (string, []string) {
	return slice[0], append(slice[1:])
}

func main() {
	inputFile, _ := ioutil.ReadFile("input.txt")
	tokens := tokenize(string(inputFile))
	cleanTokens, garbageRemoved := removeGarbage(removeBangs(tokens))
	fmt.Println(cleanTokens)
	fmt.Println(calcScore(cleanTokens))
	fmt.Println(garbageRemoved)
}
