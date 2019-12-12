package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sy, sx, gy, gx, c := handleInput()
	seen := make([][]bool, len(c))
	for i := range seen {
		seen[i] = make([]bool, len(c[0]))
	}
	// fmt.Println(sy, sx, gy, gx, seen)
	if dfs(sy, sx, gy, gx, c, seen) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func dfs(y, x, gy, gx int, c [][]string, seen [][]bool) bool {
	//chech if out of range
	if y < 0 || x < 0 || len(c)-1 < y || len(c[0])-1 < x || c[y][x] == "#" {
		return false
	}
	//goal
	if y == gy && x == gx {
		return true
	}
	//already walked
	if seen[y][x] {
		return false
	}

	//mark
	seen[y][x] = true

	//up
	if dfs(y-1, x, gy, gx, c, seen) {
		return true
	}
	//down
	if dfs(y+1, x, gy, gx, c, seen) {
		return true
	}
	//left
	if dfs(y, x-1, gy, gx, c, seen) {
		return true
	}
	//right
	if dfs(y, x+1, gy, gx, c, seen) {
		return true
	}
	return false
}

//handleInput returns y of s, x of s, y of g, x of g and map which is generated from stdin.
func handleInput() (int, int, int, int, [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	splited := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(splited[0])
	c := [][]string{}
	sx := 0
	sy := 0
	gx := 0
	gy := 0
	for i := 0; i < h; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), "")
		for sIdx, v := range row {
			if v == "s" {
				sx = sIdx
				sy = i
			}
		}
		for gIdx, v := range row {
			if v == "g" {
				gx = gIdx
				gy = i
			}
		}

		c = append(c, row)
	}
	return sy, sx, gy, gx, c
}
