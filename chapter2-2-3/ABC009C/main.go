package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	totalChangeCnt, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

	// var sortedCArrau []char
	scanner.Scan()
	for i, v := range scanner.Text() {
		cArray = append(cArray, char{v, false, i, i})
		// sortedCArrau = append(sortedCArrau, char{v, false})
	}

	changeCnt := 0
	strLen := len(cArray)
	for i := 0; i < strLen; i++ {
		if totalChangeCnt <= changeCnt {
			break
		}

		currentMinIdx := getMinIdxFrom(i)
		if currentMinIdx == i {
			//do nothing
			continue
		}
		//swap
		if totalChangeCnt-changeCnt == 1 {
			//char at i was once moved
			if cArray[i].movedOnce {
				if !cArray[currentMinIdx].movedOnce {
					changeCnt++
					cArray[currentMinIdx].movedOnce = true
				}
				cArray[i], cArray[currentMinIdx] = cArray[currentMinIdx], cArray[i]
				continue
			}
			//char at i is never moved
			//find min char which was moved once
			onceMovedMin := getMovedMinIdxFrom(i)
			// fmt.Println(onceMovedMin)
			if onceMovedMin == i {
				//There is nothing to do
				break
			}
			//swap
			cArray[i], cArray[onceMovedMin] = cArray[onceMovedMin], cArray[i]
			changeCnt++
			break
		}

		if !cArray[currentMinIdx].movedOnce {
			changeCnt++
			cArray[currentMinIdx].movedOnce = true
		}
		if !cArray[i].movedOnce {
			changeCnt++
			cArray[i].movedOnce = true
		}
		cArray[i], cArray[currentMinIdx] = cArray[currentMinIdx], cArray[i]
	}

	for _, v := range cArray {
		fmt.Print(string(v.c))
	}
	fmt.Println()
}

func getMinIdxFrom(start int) int {
	minIdx := start
	for i := start; i < len(cArray); i++ {
		if cArray[i].c <= cArray[minIdx].c {
			minIdx = i
		}
	}
	return minIdx
}

func getMovedMinIdxFrom(start int) int {
	minIdx := start
	for i := start; i < len(cArray); i++ {
		if cArray[i].c <= cArray[minIdx].c && cArray[i].movedOnce {
			minIdx = i
		}
	}
	return minIdx
}
