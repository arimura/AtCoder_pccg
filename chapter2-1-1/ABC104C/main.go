package main

import "fmt"

type inputRec struct {
	//100, 200 or ...
	point int
	//number of problem
	problemNum int
	//bonus
	bonus int
}

type problem struct {
	//100, 200 or ...
	point int
	//solved
	solved bool
	//next problem
	next *problem
}

func main() {
	//D
	// d := 2
	// g := 700
	recs := make([]inputRec, 2)
	recs[0] = inputRec{
		100,
		3,
		500,
	}
	recs[1] = inputRec{
		200,
		5,
		800,
	}
	// totalNumOfProblem := 0
	var problems *problem
	var tmp *problem
	for outerIdx, inputRec := range recs {
		for i := 0; i < inputRec.problemNum; i++ {
			p := problem{
				inputRec.point,
				false,
				nil,
			}
			if !(outerIdx == 0 && i == 0) {
				tmp.next = &p
			} else {
				problems = &p
			}
			tmp = &p
		}
	}

	// dump(problems)
	print(problems)
}

func permutation(problems *problem) {

}

func print(problems *problem) {
	fmt.Printf("(%d, %t), ", problems.point, problems.solved)
	if problems.next != nil {
		print(problems.next)
	}
}

// func dump(problem *problem) {
// 	fmt.Println(problem)
// 	if problem.next != nil {
// 		dump(problem.next)
// 	}
// }
