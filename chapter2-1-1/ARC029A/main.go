package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	length, tArray := handleInput()
	// fmt.Println(length)
	// fmt.Println(tArray)
	min := math.MaxInt32
	for bit := 0; bit < 1<<length; bit++ {
		b1, b2 := assign(bit, tArray)
		// fmt.Printf("b1: %v, b1: %v\n", b1, b2)
		t := calcTime(b1, b2)
		if t <= min {
			min = t
		}
	}
	fmt.Println(min)
}

func calcTime(b1, b2 []int) int {
	t1 := 0
	t2 := 0
	for _, t := range b1 {
		t1 += t
	}
	for _, t := range b2 {
		t2 += t
	}
	if t1 <= t2 {
		return t2
	}
	return t1
}

func assign(bit int, tArray []int) ([]int, []int) {
	var burner1 []int
	var burner2 []int
	for i, t := range tArray {
		if bit&(1<<uint(i)) == 0 {
			burner1 = append(burner1, t)
		} else {
			burner2 = append(burner2, t)
		}
	}
	return burner1, burner2
}

func handleInput() (uint, []int) {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	l, _ := strconv.Atoi(stdin.Text())
	u := uint(l)
	tArray := make([]int, l)
	for i := 0; i < l; i++ {
		stdin.Scan()
		tArray[i], _ = strconv.Atoi(stdin.Text())
	}
	return u, tArray
}
