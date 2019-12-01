package main

import (
	"fmt"
	"math"
)

func main() {
	// stdin := bufio.NewScanner(os.Stdin)
	// stdin.Scan()
	// in := stdin.Text()
	// fmt.Print(in)
	combinatin([]int{1, 2, 4}, []bool{false, false, false}, 0)

	fmt.Println(createNum([]int{1, 2, 3}))
}

func combinatin(array []int, plusArray []bool, i int) {
	//last element
	if len(plusArray)-1 == i {
		//dump
		fmt.Println(plusArray)
		//change
		plusArray[i] = true
		//dump
		fmt.Println(plusArray)
		//back to false
		plusArray[i] = false
		return
	}
	//now false
	combinatin(array, plusArray, i+1)
	//upward
	plusArray[i] = true
	combinatin(array, plusArray, i+1)
	plusArray[i] = false
}

// func calcSum(array []int, plusArray []bool, plusIndex int) int {
// 	if plusIndex == len(plusArray)-1 {

// 	}
// }

func createNum(array []int) int {
	sum := 0
	digit := 0
	for i := len(array) - 1; 0 <= i; i-- {
		sum += int(math.Pow10(digit)) * array[i]
		digit++
	}
	return sum
}
