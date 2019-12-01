package main

import (
	"fmt"
)

func main() {
	// stdin := bufio.NewScanner(os.Stdin)
	// stdin.Scan()
	// in := stdin.Text()
	// fmt.Print(in)
	combinatin([]int{1, 2, 4}, []bool{false, false, false}, 0)
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

// func calcSum(array []int, plusArray []bool) int {

// }
