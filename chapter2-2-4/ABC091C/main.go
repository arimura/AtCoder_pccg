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
	x, y int
}

type pairCandidates struct {
	red           point
	blueCandidate map[point]bool
}

type byNumberOfCandidate []pairCandidates

func (a byNumberOfCandidate) Len() int {
	return len(a)
}

func (a byNumberOfCandidate) Less(i, j int) bool {
	return len(a[i].blueCandidate) < len(a[j].blueCandidate)
}

func (a byNumberOfCandidate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	redPoints := []point{}
	bluePoints := []point{}
	idx := 0
	for scanner.Scan() {
		pointArray := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(pointArray[0])
		y, _ := strconv.Atoi(pointArray[1])
		if idx < l {
			redPoints = append(redPoints, point{x, y})
		} else {
			bluePoints = append(bluePoints, point{x, y})
		}
		idx++
	}

	blueCandidatesForRed := []pairCandidates{}
	for _, r := range redPoints {
		blueCandidates := map[point]bool{}
		for _, b := range bluePoints {
			if r.x < b.x && r.y < b.y {
				blueCandidates[b] = true
			}
		}
		blueCandidatesForRed = append(blueCandidatesForRed, pairCandidates{r, blueCandidates})
	}

	//sort
	sort.Sort(byNumberOfCandidate(blueCandidatesForRed))

	//reduce
	pairCnt := 0
	for {
		fmt.Println(blueCandidatesForRed)
		if len(blueCandidatesForRed) == 0 {
			break
		}

		pairCandidate := blueCandidatesForRed[0]
		blueCandidates := pairCandidate.blueCandidate
		if len(blueCandidates) == 0 {
			blueCandidatesForRed = blueCandidatesForRed[1:]
			continue
		}
		var fixedBlue point
		for p := range blueCandidates {
			fixedBlue = p
			break
		}
		pairCnt++
		blueCandidatesForRed = blueCandidatesForRed[1:]
		//remove fixedBlue from blueCandidates
		for _, nextPairCandidate := range blueCandidatesForRed {
			delete(nextPairCandidate.blueCandidate, fixedBlue)
		}

		sort.Sort(byNumberOfCandidate(blueCandidatesForRed))
	}
	fmt.Println(pairCnt)
}
