package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	y int
	x int
}

func main() {
	landMap, seaAreas, startY, startX := handleInput()

	for _, seaArea := range seaAreas {
		landMap[seaArea.y][seaArea.x] = 'o'
		var landChecked [][]bool
		for _, row := range landMap {
			landChecked = append(landChecked, make([]bool, len(row)))
		}
		dfs(landMap, &landChecked, startY, startX)
		if isAllChecked(landMap, landChecked) {
			fmt.Println("YES")
			return
		}
		landMap[seaArea.y][seaArea.x] = 'x'
	}
	fmt.Println("NO")
}

func isAllChecked(landMap [][]rune, landChecked [][]bool) bool {
	for y, row := range landMap {
		for x, v := range row {
			if v == 'o' && !landChecked[y][x] {
				return false
			}
		}
	}
	return true
}

func dfs(landMap [][]rune, landChecked *[][]bool, y, x int) {
	if y < 0 || x < 0 || len(landMap)-1 < y || len(landMap[0])-1 < x || landMap[y][x] == 'x' || (*landChecked)[y][x] {
		return
	}
	(*landChecked)[y][x] = true
	dfs(landMap, landChecked, y-1, x)
	dfs(landMap, landChecked, y+1, x)
	dfs(landMap, landChecked, y, x-1)
	dfs(landMap, landChecked, y, x+1)
}

func copy(src [][]bool) [][]bool {
	copied := make([][]bool, 0)
	for _, row := range src {
		copiedRow := make([]bool, 0)
		for _, v := range row {
			copiedRow = append(copiedRow, v)
		}
		copied = append(copied, copiedRow)
	}
	return copied
}

// handleInput return 2 dimensions rune([y][x]rune) and slice of sea areas
func handleInput() ([][]rune, []point, int, int) {
	var landMap [][]rune
	var searAreas []point
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	startY := 0
	startX := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		for x, r := range runes {
			if r == 'x' {
				searAreas = append(searAreas, point{y, x})
			}
			if r == 'o' {
				startY = y
				startX = x
			}
		}
		landMap = append(landMap, runes)
		y++
	}
	return landMap, searAreas, startY, startX
}
