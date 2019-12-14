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
	landMap, _ := handleInput()

	var landChecked [][]bool
	for _, row := range landMap {
		landChecked = append(landChecked, make([]bool, len(row)))
	}

	var x, y int
	for yIdx, row := range landMap {
		for xIdx, v := range row {
			if v == 'o' {
				y = yIdx
				x = xIdx
				break
			}
		}
	}

	f := false
	if dfs(landMap, &landChecked, y, x, &f) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func uncheckedLandNum(landMap [][]rune, landChecked *[][]bool) int {
	n := 0
	for y, row := range landMap {
		for x, v := range row {
			if v == 'o' && !(*landChecked)[y][x] {
				n++
			}
		}
	}
	return n
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

func dfs(landMap [][]rune, landChecked *[][]bool, y, x int, useFilling *bool) bool {
	//check unchecked land exists
	n := uncheckedLandNum(landMap, landChecked)
	if n == 0 {
		return true
	}

	if y < 0 || x < 0 || len(landMap)-1 < y || len(landMap[0])-1 < x || (*landChecked)[y][x] {
		//out of range or alreadu checked
		return false
	}
	//use filling
	if landMap[y][x] == 'x' && !*useFilling {
		fmt.Printf("on y:%d, x:%d by fullfilling\n", y, x)

		copied := copy(*landChecked)

		(*landChecked)[y][x] = true
		*useFilling = true
		if dfs(landMap, &copied, y+1, x, useFilling) {
			return true
		}
		if dfs(landMap, &copied, y-1, x, useFilling) {
			return true
		}
		if dfs(landMap, &copied, y, x-1, useFilling) {
			return true
		}
		if dfs(landMap, &copied, y, x+1, useFilling) {
			return true
		}
		*useFilling = false
		(*landChecked)[y][x] = false
		return false
	}

	if landMap[y][x] == 'x' {
		//do nothing
		return false
	}

	//on land
	fmt.Printf("on y:%d, x:%d\n", y, x)
	(*landChecked)[y][x] = true
	if dfs(landMap, landChecked, y+1, x, useFilling) {
		return true
	}
	if dfs(landMap, landChecked, y-1, x, useFilling) {
		return true
	}
	if dfs(landMap, landChecked, y, x-1, useFilling) {
		return true
	}
	if dfs(landMap, landChecked, y, x+1, useFilling) {
		return true
	}
	return false
}

// handleInput return 2 dimensions rune([y][x]rune) .
func handleInput() ([][]rune, []point) {
	var landMap [][]rune
	var lands []point
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		for x, r := range runes {
			if r == 'o' {
				lands = append(lands, point{y, x})
			}
		}
		landMap = append(landMap, runes)
		y++
	}
	return landMap, lands
}
