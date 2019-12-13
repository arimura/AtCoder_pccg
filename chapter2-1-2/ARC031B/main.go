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
	// fmt.Println())
	landMap, lands := handleInput()

	var landChecked [][]bool
	for _, row := range landMap {
		landChecked = append(landChecked, make([]bool, len(row)))
	}

	//dequeue
	land := lands[0]
	lands = lands[1:]

	var continuousLand []point
	dfs(landMap, landChecked, &lands, &continuousLand, land.y, land.x)
	fmt.Printf("lands: %v\n", lands)
	fmt.Printf("continuousLand: %v\n", continuousLand)

	land = dequeueAdjecentLand(&landMap, &lands, &continuousLand)
	fmt.Printf("lands: %v\n", land)
	dfs(landMap, landChecked, &lands, &continuousLand, land.y, land.x)
	fmt.Printf("lands: %v\n", lands)
	fmt.Printf("continuousLand: %v\n", continuousLand)

	if len(lands) == 0 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func dequeueAdjecentLand(landMap *[][]rune, lands, continuousLand *[]point) point {
	var f point
	for _, partOfContinuousLand := range *continuousLand {
		for i, land := range *lands {
			//check land is arouncd partOfContinuousLand
			y := partOfContinuousLand.y - land.y
			if y < 0 {
				y = -y
			}
			x := partOfContinuousLand.x - land.x
			if x < 0 {
				x = -x
			}

			if y+x == 2 {
				//found
				f = (*lands)[i]
				*lands = append((*lands)[:i], (*lands)[i+1:]...)

				(*landMap)[partOfContinuousLand.y-((partOfContinuousLand.y-land.y)/2)][partOfContinuousLand.x-((partOfContinuousLand.x-land.x)/2)] = 'o'
				fmt.Printf("new mapped x: %d, y %d", partOfContinuousLand.x-((partOfContinuousLand.x-land.x)/2), partOfContinuousLand.y-((partOfContinuousLand.y-land.y)/2))
				return f
			}
		}
	}
	return f
}

func dfs(landMap [][]rune, landChecked [][]bool, lands, continuousLand *[]point, y, x int) {
	if y < 0 || x < 0 || len(landMap)-1 < y || len(landMap[0])-1 < x || landMap[y][x] == 'x' || landChecked[y][x] {
		//no land
		return
	}

	landChecked[y][x] = true

	//found land.
	//remove from lands
	for i, land := range *lands {
		if land.y == y && land.x == x {
			*lands = append((*lands)[:i], (*lands)[i+1:]...)
			break
		}
	}
	//add to continuousLand
	*continuousLand = append(*continuousLand, point{y, x})

	//up, down, left, rigth
	dfs(landMap, landChecked, lands, continuousLand, y+1, x)
	dfs(landMap, landChecked, lands, continuousLand, y-1, x)
	dfs(landMap, landChecked, lands, continuousLand, y, x-1)
	dfs(landMap, landChecked, lands, continuousLand, y, x+1)
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
