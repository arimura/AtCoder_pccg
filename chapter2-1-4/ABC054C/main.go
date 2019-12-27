package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seen = [9]bool{}
var n = 0
var vertices = [9][]int{}
var completeCnt = 0

func main() {
	handleInput()
	perm(1, 0)
	fmt.Println(completeCnt)
}

func perm(id, cnt int) {
	if seen[id] {
		//already seen
		return
	}
	//first visit
	seen[id] = true
	cnt++

	//check cnt
	if n <= cnt {
		//perm complete
		completeCnt++
		seen[id] = false
		return
	}

	//next
	for _, adj := range vertices[id] {
		perm(adj, cnt)
	}

	seen[id] = false
}

func handleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	h := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(h[0])
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(edge[0])
		r, _ := strconv.Atoi(edge[1])
		vertices[l] = append(vertices[l], r)
		vertices[r] = append(vertices[r], l)
	}
}
