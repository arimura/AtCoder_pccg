package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()
	if a == "a" {
		fmt.Println(-1)
		return
	}
	fmt.Println("a")
}
