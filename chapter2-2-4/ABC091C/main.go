package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
	next *[]*edge
}

func (a point) hasNext() bool {
	for _, edge := range *a.next {
		if !edge.used {
			return true
		}
	}
	return false
}

type edge struct {
	used bool
	to   *point
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	sPoint := &point{-1, -1, &[]*edge{}}
	tPoint := &point{-2, -2, &[]*edge{}}

	redPoints := []*point{}
	idx := 0
	for scanner.Scan() {
		pointArray := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(pointArray[0])
		y, _ := strconv.Atoi(pointArray[1])
		if idx < l {
			r := point{x, y, &[]*edge{}}
			*sPoint.next = append(*(sPoint.next), &edge{false, &r})
			redPoints = append(redPoints, &r)
		} else {
			b := &point{x, y, &[]*edge{&edge{false, tPoint}}}
			for _, r := range redPoints {
				if r.x < x && r.y < y {
					*r.next = append(*r.next, &edge{false, b})
				}
			}
		}
		idx++
	}

	cnt := 0
	for _, sNext := range *sPoint.next {
		if sNext.used {
			continue
		}
		if findPathToT(sNext.to) {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func findPathToT(current *point) bool {
	if current.x == -2 && current.y == -2 {
		return true
	}

	if !current.hasNext() {
		return false
	}

	for _, next := range *current.next {
		if next.used {
			continue
		}
		next.used = true
		arrived := findPathToT(next.to)
		if arrived {
			//残余
			*next.to.next = append(*next.to.next, &edge{false, current})
			return true
		}
		next.used = false
	}
	return false
}
