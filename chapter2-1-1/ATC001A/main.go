package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	y, x, c := handleInput()
	stack := make([]string, 0)
	if dfs(y, x, c, &stack) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func push(y, x int, stack *[]string) {
	v := strconv.Itoa(y) + "-" + strconv.Itoa(x)
	*stack = append(*stack, v)
}

func pop(stack *[]string) {
	*stack = (*stack)[:len(*stack)-1]
}

func inStack(y, x int, stack []string) bool {
	v := strconv.Itoa(y) + "-" + strconv.Itoa(x)
	for _, e := range stack {
		if e == v {
			return true
		}
	}
	return false
}

func dfs(y, x int, c [][]string, stack *[]string) bool {
	if c[y][x] == "g" {
		return true
	}
	push(y, x, stack)

	//go up
	if 0 < y && c[y-1][x] != "#" && !inStack(y-1, x, *stack) {
		if dfs(y-1, x, c, stack) {
			return true
		}
	}
	//go down
	if y < len(c)-1 && c[y+1][x] != "#" && !inStack(y+1, x, *stack) {
		if dfs(y+1, x, c, stack) {
			return true
		}
	}
	//go left
	if 0 < x && c[y][x-1] != "#" && !inStack(y, x-1, *stack) {
		if dfs(y, x-1, c, stack) {
			return true
		}
	}
	//go right
	if x < len(c[0])-1 && c[y][x+1] != "#" && !inStack(y, x+1, *stack) {
		if dfs(y, x+1, c, stack) {
			return true
		}
	}

	pop(stack)
	return false
}

// retrun [y][x]string
func handleInput() (int, int, [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	splited := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(splited[0])
	c := [][]string{}
	x := 0
	y := 0
	for i := 0; i < h; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), "")
		for sIdx, v := range row {
			if v == "s" {
				x = sIdx
				y = i
			}
		}

		c = append(c, row)
	}
	return y, x, c
}
