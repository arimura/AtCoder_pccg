package main

import (
	"bufio"
	"os"
)

type point struct {
	y int
	x int
}

func main() {
	// fmt.Println())
	landMap, lands := handleInput()
	land := lands[1:]
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
