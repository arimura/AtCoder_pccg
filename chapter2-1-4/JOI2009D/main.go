package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	picked           = [11]bool{}
	generatedNumbers = map[string]bool{}
	numbers          = []string{}
	pickNum          = 0
)

func main() {
	handleInput()
	a := make([]string, 0)
	for i := range numbers {
		perm(i, 0, &a)
	}
	fmt.Println(len(generatedNumbers))
}

func perm(idx, cnt int, array *[]string) {
	if picked[idx] {
		return
	}

	//fitst pick
	cnt++
	picked[idx] = true
	*array = append(*array, numbers[idx])

	if pickNum <= cnt {
		generatedNumbers[strings.Join(*array, "")] = true
		picked[idx] = false
		*array = (*array)[:len(*array)-1]
		return
	}

	//next
	for i := range numbers {
		perm(i, cnt, array)
	}
	picked[idx] = false
	*array = (*array)[:len(*array)-1]
}

func handleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	pickNum, _ = strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
}
