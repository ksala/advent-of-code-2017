package main

import (
	"fmt"
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

func main() {
	input := make([]int, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}
	list := makeCircularList(append([]int{}, input...))
	lenghts := []int{130, 126, 1, 11, 140, 2, 255, 207, 18, 254, 246, 164, 29, 104, 0, 224}
	for _, lenght := range lenghts {
		elems := list.getElems(lenght)
		rev := reverseSlice(elems)
		list.setElems(lenght, rev)
		list.moveCur(lenght + list.skip)
		list.skip = list.skip + 1
	}
	fmt.Printf("First answer is %d\n", list.elems[0]*list.elems[1])

	list2 := makeCircularList(append([]int{}, input...))
	lenghtsInput2 := "130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224"
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
	fmt.Printf("The second answer is: ")
	for _, value := range denseHash {
		fmt.Printf("%.2x", value)
	}
	fmt.Println()
}
