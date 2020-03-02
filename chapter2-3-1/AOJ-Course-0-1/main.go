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
	memo   [100][10000]int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(t[0])
	totalW, _ = strconv.Atoi(t[1])

	i := 0
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " ")
		v[i], _ = strconv.Atoi(t[0])
		w[i], _ = strconv.Atoi(t[1])
		i++
	}

	maxV := 0
	un := uint(n)
	for bit := 0; bit < (1 << un); bit++ {
		fmt.Printf("bit: %04b\n", bit)
		tmpW := 0
		tmpV := 0
		for d := n - 1; 0 <= d; d-- {
			if 0 < bit&(1<<uint(d)) {
				if tmpW+w[d] <= totalW {
					tmpW += w[d]
					tmpV += v[d]
				}
			}
		}
		fmt.Println(tmpV)
		if maxV < tmpV {
			maxV = tmpV
		}
	}
	fmt.Println(maxV)
}
