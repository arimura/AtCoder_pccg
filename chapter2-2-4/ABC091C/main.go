package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y    int
	edges   *[]*edge
	fromNum int
}

func (a point) hasNext() bool {
	for _, edge := range *a.edges {
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

type redSorter []*edge

func (a redSorter) Len() int { return len(a) }
func (a redSorter) Less(i, j int) bool {
	return len(*a[i].to.edges) < len(*a[j].to.edges)
}
func (a redSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type blueSorter []*edge

func (a blueSorter) Len() int { return len(a) }
func (a blueSorter) Less(i, j int) bool {
	return a[i].to.fromNum < a[j].to.fromNum
}
func (a blueSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	sPoint := &point{-1, -1, &[]*edge{}, 0}
	tPoint := &point{-2, -2, &[]*edge{}, 0}

	redPoints := []*point{}
	idx := 0
	for scanner.Scan() {
		pointArray := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(pointArray[0])
		y, _ := strconv.Atoi(pointArray[1])
		if idx < l {
			r := point{x, y, &[]*edge{}, 0}
			*sPoint.edges = append(*(sPoint.edges), &edge{false, &r})
			redPoints = append(redPoints, &r)
		} else {
			b := &point{x, y, &[]*edge{&edge{false, tPoint}}, 0}
			for _, r := range redPoints {
				if r.x < x && r.y < y {
					*r.edges = append(*r.edges, &edge{false, b})
					b.fromNum++
				}
			}
		}
		idx++
	}

	sort.Sort(redSorter(*sPoint.edges))
	for _, edge := range *sPoint.edges {
		sort.Sort(blueSorter(*edge.to.edges))
	}

	cnt := 0
	for _, sNext := range *sPoint.edges {
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
	// fmt.Println(current)
	if current.x == -2 && current.y == -2 {
		return true
	}

	if !current.hasNext() {
		return false
	}

	for _, next := range *current.edges {
		if next.used {
			continue
		}
		next.used = true
		arrived := findPathToT(next.to)
		if arrived {
			//残余
			*next.to.edges = append(*next.to.edges, &edge{false, current})
			return true
		}
		next.used = false
	}
	return false
}
