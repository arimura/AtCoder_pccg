package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	totalW int
	n      int
	v      [1000]int
	w      [1000]int
	memo   [1000][100000]int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(t[0])
	totalW, _ = strconv.Atoi(t[1])

	i := 0
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " ")
		v[i], _ = strconv.Atoi(t[0])
		w[i], _ = strconv.Atoi(t[1])
		i++
	}
	// maxV := 0
	fmt.Println(search(0, totalW))
}

func search(i, u int) int {
	if memo[i][u] != 0 {
		return memo[i][u]
	}

	if i == n {
		return 0
	} else if u < w[i] {
		return search(i+1, u)
	} else {
		r1 := search(i+1, u)

		r2 := search(i+1, u-w[i]) + v[i]
		if r1 < r2 {
			memo[i][u] = r2
			return r2
		}
		memo[i][u] = r1
		return r1
	}
}
