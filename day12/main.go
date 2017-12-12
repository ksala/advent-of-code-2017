package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type pipe map[string][]string

func appendIfMissing(slice []string, s string) []string {
	for _, value := range slice {
		if value == s {
			return slice
		}
	}
	return append(slice, s)
}

func main() {
	inputFile, _ := ioutil.ReadFile("input.txt")

	pipes := pipe{}

	for _, line := range strings.Split(string(inputFile), "\n") {
		line := strings.Split(line, " ")
		program := line[0]
		connections := line[2:]
		connections = strings.Split(strings.Join(connections, ""), ",")
		pipes[program] = append(pipes[program], connections...)
	}

	inGroup := []string{"0"}
	previousLen := 0
	groups := 0
	
	for len(pipes) > 0 {
		for {
			previousLen = len(inGroup)
			for node, connections := range pipes {
				for _, value := range inGroup {
					if node == value {
						for _, v := range connections {
							inGroup = appendIfMissing(inGroup, v)
						}
					}
				}
			}
			// We didn't append anything new in this cycle
			// So we can break out
			if len(inGroup) == previousLen {
				break
			}
		}
		if inGroup[0] == "0" {
			fmt.Printf("%d Elements talk to 0\n", len(inGroup))
		}
		for _, item := range inGroup {
			delete(pipes, item)
		}
		groups++
		for _, v := range pipes {
			inGroup = v
			break
		}
	}
	fmt.Printf("There are %d groups\n", groups)
}