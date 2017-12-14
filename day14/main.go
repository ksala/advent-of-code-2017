package main

import (
	"fmt"
	"strconv"
)

type circularList struct {
	elems []int
	cur   int
	skip  int
}

func (c circularList) getElems(n int) []int {
	if n <= (len(c.elems) - c.cur) {
		return c.elems[c.cur : c.cur+n]
	}
	return append(c.elems[c.cur:], c.elems[:n+c.cur-len(c.elems)]...)
}

func (c *circularList) setElems(n int, put []int) {
	elementsToEnd := len(c.elems) - c.cur
	begin := []int{}
	middle := []int{}
	end := []int{}
	if len(put) < elementsToEnd {
		begin = c.elems[:c.cur]
		middle = put
		end = c.elems[len(put)+c.cur:]
	} else {
		begin = put[elementsToEnd:]
		middle = c.elems[len(put)-elementsToEnd : c.cur]
		end = put[:elementsToEnd]
	}
	t := append(middle, end...)
	t = append(begin, t...)
	c.elems = t
}

func (c *circularList) moveCur(n int) {
	c.cur = (c.cur + n) % len(c.elems)
}

func makeCircularList(elems []int) *circularList {
	list := &circularList{elems: elems}
	return list
}

func reverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	copy(reversed, slice)
	for i := len(reversed)/2 - 1; i >= 0; i-- {
		opp := len(reversed) - 1 - i
		reversed[i], reversed[opp] = reversed[opp], reversed[i]
	}
	return reversed
}

func markVisited(rowIndex, cellIndex int, grid [][]bool) {
	if rowIndex < 0 || rowIndex >= len(grid) {
		return
	}
	if cellIndex < 0 || cellIndex >= len(grid[rowIndex]) {
		return
	}
	if !grid[rowIndex][cellIndex] {
		return
	}

	grid[rowIndex][cellIndex] = false
	markVisited(rowIndex+1, cellIndex, grid)
	markVisited(rowIndex-1, cellIndex, grid)
	markVisited(rowIndex, cellIndex+1, grid)
	markVisited(rowIndex, cellIndex-1, grid)
}

func main() {
	input := make([]int, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}

	var sum int
	grid := make([][]bool, 128)

	for i := 0; i < 128; i++ {
		grid[i] = make([]bool, 128)
		list2 := makeCircularList(append([]int{}, input...))
		lenghtsInput2 := fmt.Sprintf("uugsqrei-%d", i)
		byteArray := []byte(lenghtsInput2)
		byteArray = append(byteArray, 17, 31, 73, 47, 23)
		lenghts2 := make([]int, len(byteArray))
		for index, value := range byteArray {
			lenghts2[index] = int(value)
		}

		for round := 0; round < 64; round++ {
			for _, lenght := range lenghts2 {
				elems := list2.getElems(lenght)
				rev := reverseSlice(elems)
				list2.setElems(lenght, rev)
				list2.moveCur(lenght + list2.skip)
				list2.skip = list2.skip + 1
			}
		}
		denseHash := []int{}
		for i := 0; i < 256; i = i + 16 {
			total := 0
			for _, elem := range list2.elems[i : i+16] {
				total = total ^ elem
			}
			denseHash = append(denseHash, total)
		}
		var hash string
		for _, value := range denseHash {
			hash = fmt.Sprintf("%s%.2x", hash, value)
		}

		for indexRow, b := range hash {
			conv, _ := strconv.ParseInt(string(b), 16, 8)
			bits := fmt.Sprintf("%04b", conv)
			for indexBit, bit := range bits {
				if bit == '1' {
					sum++
					grid[i][indexRow*4+indexBit] = true
				}
			}
		}
	}
	var groups int
	for rowIndex, row := range grid {
		for cellIndex, cell := range row {
			if cell {
				groups++
				markVisited(rowIndex, cellIndex, grid)
			}
		}
	}
	fmt.Printf("Used squares: %d\n", sum)
	fmt.Printf("Groups: %d\n", groups)
}
