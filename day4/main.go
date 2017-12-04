package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func orderWord(word string) string {
	slicedWord := strings.Split(word, "")
	sort.Strings(slicedWord)
	return strings.Join(slicedWord, "")
}

func validPassphrase(passphrase string, level int) bool {
	passphraseSlice := strings.Split(passphrase, " ")
	for {
		if len(passphraseSlice) == 0 {
			break
		}
		wordCheck := passphraseSlice[0]
		orderedWordCheck := orderWord(wordCheck)
		for _, word := range passphraseSlice[1:] {
			if word == wordCheck {
				return false
			}
			if level == 2 {
				if orderedWordCheck == orderWord(word) {
					return false
				}
			}
		}
		passphraseSlice = passphraseSlice[1:]
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// Part 1
	scanner := bufio.NewScanner(file)
	totalPart1 := 0
	totalPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if validPassphrase(line, 1) {
			totalPart1++
		}
		if validPassphrase(line, 2) {
			totalPart2++
		}
	}
	fmt.Println(totalPart1)
	fmt.Println(totalPart2)

}
