package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	cArray []char
)

type char struct {
	c          rune
	movedOnce  bool
	orgIdx     int
	currentIdx int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// totalChangeCnt, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

	// var sortedCArrau []char
	scanner.Scan()
	for i, v := range scanner.Text() {
		cArray = append(cArray, char{v, false, i, i})
		// sortedCArrau = append(sortedCArrau, char{v, false})
	}
	fmt.Println(cArray)
	// fmt.Println(getMinIdxFrom(3))
}

func getMinIdxFrom(start int) int {
	minIdx := start
	for i := start; i < len(cArray); i++ {
		if cArray[i].c < cArray[minIdx].c {
			minIdx = i
		}
	}
	return minIdx
}
