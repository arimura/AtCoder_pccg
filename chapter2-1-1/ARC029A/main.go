package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print(handleInput())
}

func handleInput() (int, []int) {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	l, _ := strconv.Atoi(stdin.Text())
	tArray := make([]int, l)
	for i := 0; i < l; i++ {
		stdin.Scan()
		tArray[i], _ = strconv.Atoi(stdin.Text())
	}
	return l, tArray
}
