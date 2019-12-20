package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	y, x int
}

var m = [1002][1002]rune{}
var timeMap = [1002][1002]int{}
var q = []point{}
var directions = []point{point{1, 0}, point{0, 1}, point{-1, 0}, point{0, -1}}
var time = 0

func main() {
	n := handleInput()
	nextCheese := '1'
	t := 0
	s := q[0]

	for len(q) != 0 {
		//push
		p := q[0]
		q = q[1:]

		//check if cheese
		if m[p.y][p.x] == nextCheese {
			//got cheese. record time
			t += timeMap[p.y][p.x]
			if int(nextCheese)-48 == n {
				break
			}
			nextCheese++
			//Initilize state
			timeMap = [1002][1002]int{}
			q = []point{p}
			s = p
			continue
		}

		//go to next point
		for _, d := range directions {
			n := point{p.y + d.y, p.x + d.x}
			//no X AND no '0' AND first time
			if m[n.y][n.x] != 'X' && m[n.y][n.x] != 0 && timeMap[n.y][n.x] == 0 && n != s {
				q = append(q, n)
				timeMap[n.y][n.x] = timeMap[p.y][p.x] + 1
			}
		}
	}
	fmt.Println(t)
}

func handleInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[2])
	y := 1
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			m[y][x+1] = r
			if r == 'S' {
				q = append(q, point{y, x + 1})
			}
		}
		y++
	}
	return n
}
