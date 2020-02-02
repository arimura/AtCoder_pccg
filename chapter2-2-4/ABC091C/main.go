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
	next *[]*point
	prev *[]*point
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	sPoint := &point{-1, -1, &[]*point{}, &[]*point{}}
	tPoint := &point{-2, -2, &[]*point{}, &[]*point{}}

	redPoints := []*point{}
	idx := 0
	for scanner.Scan() {
		pointArray := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(pointArray[0])
		y, _ := strconv.Atoi(pointArray[1])
		if idx < l {
			r := point{x, y, &[]*point{}, &[]*point{}}
			*sPoint.next = append(*(sPoint.next), &r)
			redPoints = append(redPoints, &r)
		} else {
			b := &point{x, y, &[]*point{tPoint}, &[]*point{}}
			for _, r := range redPoints {
				if r.x < x && r.y < y {
					*r.next = append(*r.next, b)
				}
			}
		}
		idx++
	}

	cnt := 0
	current := sPoint
	for {
		fmt.Printf("current: %v, pointer: %p\n", current, &current)
		if current.x == -2 && current.y == -2 {
			fmt.Println("t")
			cnt++
			current = sPoint
			continue
		}

		fmt.Printf("next: %v\n", *current.next)
		if len(*current.next) == 0 && len(*current.prev) == 0 {
			fmt.Println("no next")
			if len(*sPoint.next) == 0 {
				break
			}
			current = sPoint
			continue
		}

		if len(*current.next) != 0 {
			prev := current
			current = (*current.next)[0]
			*prev.next = (*prev.next)[1:]
			*current.prev = append(*current.prev, prev)
			continue
		}

		if len(*current.prev) != 0 {
			prev := current
			current = (*current.prev)[0]
			*prev.prev = (*prev.prev)[1:]
			continue
		}
	}

	fmt.Println(cnt)
}
