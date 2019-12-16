package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	y, x int
}

func main() {
	m, lands, landsChecked := handleInput()
	c := 0
	for _, land := range lands {
		if landsChecked[land.y][land.x] {
			continue
		}
		dfs(m, &landsChecked, land.y, land.x)
		c++
	}
	fmt.Println(c)
}

func dfs(m [][]string, landsChecked *[][]bool, y, x int) {
	if y < 0 || x < 0 || len(m)-1 < y || len(m[0])-1 < x || (*landsChecked)[y][x] || m[y][x] == "0" {
		return
	}
	(*landsChecked)[y][x] = true
	dfs(m, landsChecked, y+1, x)
	dfs(m, landsChecked, y+1, x+1)
	dfs(m, landsChecked, y, x+1)
	dfs(m, landsChecked, y-1, x+1)
	dfs(m, landsChecked, y-1, x)
	dfs(m, landsChecked, y-1, x-1)
	dfs(m, landsChecked, y, x-1)
	dfs(m, landsChecked, y+1, x-1)
}

func handleInput() ([][]string, []point, [][]bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var m [][]string
	var lands []point
	var landsChecked [][]bool
	y := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		mapRow := make([]string, 0)
		for x, c := range row {
			mapRow = append(mapRow, c)
			if c == "1" {
				lands = append(lands, point{y, x})
			}
		}
		m = append(m, mapRow)
		landsChecked = append(landsChecked, make([]bool, len(row)))
		y++
	}
	return m, lands, landsChecked
}
