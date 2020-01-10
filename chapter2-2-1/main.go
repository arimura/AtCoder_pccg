package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	remain := 1000 - i
	cnt := 0
	conis := [6]int{500, 100, 50, 10, 5, 1}

	for {
		if remain == 0 {
			break
		}
		for _, coin := range conis {
			if coin <= remain {
				remain -= coin
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
