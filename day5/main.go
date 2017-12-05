package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func calcSteps(jumps []int, challenge string) int {
    index := 0
    steps := 0
    for {
        if index < 0 || index >= len(jumps) {
            break
        }
        steps++
        nextIndex := index + jumps[index]
        if jumps[index] < 3 || challenge == "part1" {
            jumps[index] = jumps[index] + 1
        } else {
            jumps[index] = jumps[index] - 1
        }
        index = nextIndex
    }
    return steps
}

func main() {
    inputFile, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    strSlice := strings.Split(strings.TrimSpace(string(inputFile)), "\n")
    intSlice := make([]int, len(strSlice))
    for index, value := range strSlice {
        i, err := strconv.Atoi(value)
        if err != nil {
            panic(err)
        }
        intSlice[index] = i
    }

    fmt.Println(calcSteps(append([]int(nil), intSlice...), "part1"))
    fmt.Println(calcSteps(append([]int(nil), intSlice...), "part2"))
}