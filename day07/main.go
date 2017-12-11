package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	name   string
	weight int
	childs []string
	parent string
}

func parseLine(line string) program {
	split := strings.Split(line, " ")
	name := split[0]
	weight, _ := strconv.Atoi(strings.Trim(split[1], "()"))
	up := []string(nil)
	if len(split) > 3 {
		for _, child := range split[3:] {
			up = append(up, strings.Trim(child, ","))
		}
	}
	return program{name: name, weight: weight, childs: up}
}

func updateParent(p program, parent string) program {
	return program{name: p.name, weight: p.weight, childs: p.childs, parent: parent}
}

func sumChild(p program, programs map[string]program, indentation int) int {
	if p.childs == nil {
		return p.weight
	}
	w := p.weight

	for _, child := range p.childs {
		childWeight := sumChild(programs[child], programs, indentation+4)
		w = w + childWeight
	}

	return w
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	programs := map[string]program{}
	for scanner.Scan() {
		line := scanner.Text()
		p := parseLine(line)
		programs[p.name] = p
	}
	for _, program := range programs {
		for _, child := range program.childs {
			for index, pro := range programs {
				if child == pro.name {
					programs[index] = updateParent(pro, program.name)
				}
			}
		}
	}
	base := ""
	for _, program := range programs {
		if program.parent == "" {
			base = program.name
			fmt.Printf("Base program: %s\n", program.name)
		}
	}

	s := sumChild(programs[base], programs, 0)
	fmt.Printf("%s -> %s = %d\n", programs[base].name, programs[base].childs, s)

	childs := programs[base].childs
	childSums := map[string]int{}
	for _, child := range childs {
		s := sumChild(programs[child], programs, 0)
		childSums[programs[child].name] = s
		fmt.Printf("%s -> %s = %d\n", programs[child].name, programs[child].childs, s)
	}
}
