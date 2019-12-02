package main

import (
	"bufio"
	"fmt"
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
	found := false
	combinatin(array, []bool{false, false, false}, 0, &found)
	// fmt.Println(sum([]int{1, 2, 3}, []bool{false, false}))
	// fmt.Println(sum([]int{1, 2, 3}, []bool{false, true}))
	// fmt.Println(sum([]int{1, 2, 3}, []bool{true, false}))
	// fmt.Println(sum([]int{1, 2, 3}, []bool{true, true}))
	// print([]int{1, 2, 3}, []bool{true, false})
	// print([]int{1, 2, 3}, []bool{true, true})
	// combinatin([]int{1, 2, 2, 2}, []bool{false, false, false}, 0)
	// combinatin([]int{0, 2, 9, 0}, []bool{false, false, false}, 0)
	// combinatin([]int{3, 2, 4, 2}, []bool{false, false, false}, 0)
}

func combinatin(array []int, plusArray []bool, i int, found *bool) {
	if *found {
		return
	}
	//last element
	if len(plusArray)-1 == i {
		calced := sum(array, plusArray)
		if calced == 7 {
			print(array, plusArray)
			*found = true
			return
		}

		plusArray[i] = true

		calced = sum(array, plusArray)
		if calced == 7 {
			print(array, plusArray)
			*found = true
			return
		}

		plusArray[i] = false
		return
	}
	combinatin(array, plusArray, i+1, found)
	plusArray[i] = true
	combinatin(array, plusArray, i+1, found)
	plusArray[i] = false
}

func print(array []int, opArray []bool) {
	o := ""
	for i, n := range array {
		o += strconv.Itoa(n)
		if i == len(array)-1 {
			continue
		}
		if opArray[i] {
			o += "+"
		} else {
			o += "-"
		}
	}
	o += "=7"
	fmt.Println(o)
}

// 1, 2, 3 index
//   +  -
//
func sum(array []int, opArray []bool) int {
	sum := 0
	for i, n := range array {
		if i == 0 {
			sum += array[i]
			continue
		}
		if opArray[i-1] {
			sum += n
		} else {
			sum -= n
		}
	}
	return sum
}
