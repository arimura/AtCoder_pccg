package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	id       int
	seen     bool
	adjacent []*vertex
}

func main() {
	vertices := handleInput()
	c := 0
	for _, v := range vertices {
		if v.seen {
			continue
		}
		tree := true
		dfs(&vertices, v, &tree, -1)
		if tree {
			c++
		}
	}
	fmt.Println(c)
}

//dfs returns true if vertex has tree
func dfs(vertices *[]*vertex, vertex *vertex, tree *bool, prevID int) {
	//already seen
	if vertex.seen {
		*tree = false
		return
	}
	vertex.seen = true
	//go next
	for _, n := range vertex.adjacent {
		if n.id == prevID {
			continue
		}
		dfs(vertices, n, tree, vertex.id)
	}
}

func handleInput() []*vertex {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var vertices []*vertex
	n, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	for i := 1; i <= n; i++ {
		vertices = append(vertices, &vertex{i, false, make([]*vertex, 0)})
	}
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), " ")
		vID1, _ := strconv.Atoi(edge[0])
		vID2, _ := strconv.Atoi(edge[1])
		v1 := vertices[vID1-1]
		v2 := vertices[vID2-1]
		v1.adjacent = append((*v1).adjacent, v2)
		v2.adjacent = append((*v2).adjacent, v1)
	}
	return vertices
}
