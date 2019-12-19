package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m = [52][52]rune{}
var cntMap = [51][51]int{}
var q = []point{}

type point struct {
	y, x int
}

func main() {
	s, g := handleInput()
	q = append(q, point{s.y, s.x})
	m[s.y][s.x] = '#'
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		if p.y == g.y && p.x == g.x {
			break
		}
		if m[p.y-1][p.x] == '.' {
			m[p.y-1][p.x] = '#'
			cntMap[p.y-1][p.x] = cntMap[p.y][p.x] + 1
			q = append(q, point{p.y - 1, p.x})
		}
		if m[p.y][p.x+1] == '.' {
			m[p.y][p.x+1] = '#'
			cntMap[p.y][p.x+1] = cntMap[p.y][p.x] + 1
			q = append(q, point{p.y, p.x + 1})
		}
		if m[p.y+1][p.x] == '.' {
			m[p.y+1][p.x] = '#'
			cntMap[p.y+1][p.x] = cntMap[p.y][p.x] + 1
			q = append(q, point{p.y + 1, p.x})
		}
		if m[p.y][p.x-1] == '.' {
			m[p.y][p.x-1] = '#'
			cntMap[p.y][p.x-1] = cntMap[p.y][p.x] + 1
			q = append(q, point{p.y, p.x - 1})
		}
	}
	fmt.Println(cntMap[g.y][g.x])
}

func handleInput() (point, point) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	sArray := strings.Split(scanner.Text(), " ")
	sy, _ := strconv.Atoi(sArray[0])
	sx, _ := strconv.Atoi(sArray[1])
	s := point{
		sy,
		sx,
	}
	scanner.Scan()
	gArray := strings.Split(scanner.Text(), " ")
	gy, _ := strconv.Atoi(gArray[0])
	gx, _ := strconv.Atoi(gArray[1])
	g := point{
		gy,
		gx,
	}
	y := 1
	for scanner.Scan() {
		for x, v := range scanner.Text() {
			m[y][x+1] = v
		}
		y++
	}
	return s, g
}
