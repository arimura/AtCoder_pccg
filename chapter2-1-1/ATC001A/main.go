package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	y, x, c := handleInput()
	fmt.Printf("%v, %v, %v\n", y, x, c)
	if dfs(y, x, c) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func inStack(y, x int, stack []string) bool {

}

func dfs(y, x int, c [][]string) bool {
	fmt.Printf("y:%d, x:%d\n", y, x)

	if c[y][x] == "g" {
		return true
	}
	fmt.Println("assign")
	c[y][x] = "X"

	//go up
	fmt.Println("up")
	if 0 < y && c[y-1][x] == "." {
		if dfs(y-1, x, c) {
			return true
		}
	}
	//go down
	fmt.Println("down")
	if y < len(c)-1 && c[y+1][x] == "." {
		if dfs(y+1, x, c) {
			return true
		}
	}
	//go left
	fmt.Println("left")
	if 0 < x && c[y][x-1] == "." {
		if dfs(y, x-1, c) {
			return true
		}
	}
	//go right
	fmt.Println("right")
	if x < len(c[0])-1 && c[y][x+1] == "." {
		if dfs(y, x+1, c) {
			return true
		}
	}

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
		sIdx := sort.SearchStrings(row, "s")
		if sIdx < len(row) {
			x = sIdx
			y = i
		}
		c = append(c, row)
	}
	return y, x, c
}
