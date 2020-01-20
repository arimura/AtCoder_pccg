package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputArray := strings.Split(scanner.Text(), " ")
	x, _ := strconv.Atoi(inputArray[0])
	y, _ := strconv.Atoi(inputArray[1])
	cnt := 0

	for x <= y {
		cnt++
		x = x * 2
	}
	fmt.Println(cnt)
}
