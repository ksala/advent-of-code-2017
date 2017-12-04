package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func validPassphrase(passphrase string) bool {
	passphraseSlice := strings.Split(passphrase, " ")
	for {
		if len(passphraseSlice) == 0 {
			break
		}
		wordCheck := passphraseSlice[0]
		for _, word := range passphraseSlice[1:] {
			if word == wordCheck {
				return false
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
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		if validPassphrase(scanner.Text()) {
			total = total + 1
		}
	}
	fmt.Println(total)

}
