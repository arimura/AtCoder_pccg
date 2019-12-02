package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	in := stdin.Text()
	array := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		array[i], _ = strconv.Atoi(string(in[i]))
	}
	plusArray := make([]bool, len(in)-1)

	sum := 0
	combinatin(array, plusArray, 0, &sum)
	fmt.Println(sum)
}

func combinatin(array []int, plusArray []bool, i int, sum *int) {
	//last element
	if len(plusArray)-1 == i {
		//dump
		// fmt.Println(plusArray)
		calced := calcSum(array, plusArray, 0, 0)
		// fmt.Println(calced)
		*sum += calced
		//change
		plusArray[i] = true
		//dump
		// fmt.Println(plusArray)
		calced = calcSum(array, plusArray, 0, 0)
		// fmt.Println(calced)
		*sum += calced
		//back to false
		plusArray[i] = false
		return
	}
	//now false
	combinatin(array, plusArray, i+1, sum)
	//upward
	plusArray[i] = true
	combinatin(array, plusArray, i+1, sum)
	plusArray[i] = false
}

func calcSum(array []int, plusArray []bool, plusStartIndex, plusIndex int) int {
	if plusIndex == len(plusArray) {
		return createNum(array[plusStartIndex : plusIndex+1])
	}
	if plusArray[plusIndex] {
		//plus
		return createNum(array[plusStartIndex:(plusIndex+1)]) + calcSum(array, plusArray, plusIndex+1, plusIndex+1)
	}
	return calcSum(array, plusArray, plusStartIndex, plusIndex+1)
}

func createNum(array []int) int {
	sum := 0
	digit := 0
	for i := len(array) - 1; 0 <= i; i-- {
		sum += int(math.Pow10(digit)) * array[i]
		digit++
	}
	return sum
}
