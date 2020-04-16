package main

import "fmt"

var oddJumps = map[int]int{}
var evenJumps = map[int]int{}

func oddJump(A []int, Ai int, i int) int {

	if val, ok := oddJumps[i]; ok {
		return val
	}

	min := -1
	minJ := -1
	for j := i + 1; j < len(A); j++ {
		if Ai <= A[j] {
			if min == -1 || A[j] < min {
				min = A[j]
				minJ = j
			}
		}
	}
	if minJ != -1 {
		oddJumps[i] = minJ
	}
	return minJ
}

func evenJump(A []int, Ai int, i int) int {

	if val, ok := evenJumps[i]; ok {
		return val
	}

	max := -1
	maxJ := -1
	for j := i + 1; j < len(A); j++ {
		if Ai >= A[j] {
			if A[j] > max {
				max = A[j]
				maxJ = j
			}
		}
	}
	if maxJ != -1 {
		evenJumps[i] = maxJ
	}
	return maxJ
}

func oddEvenJumpsR(A []int, i int) int {

	if len(A) == 0 {
		return 0
	}
	if len(A) == 1 {
		return 1
	}

	jumps := 1
	for i != len(A)-1 {
		res := -1
		if jumps%2 == 0 {
			res = evenJump(A, A[i], i)
		} else {
			res = oddJump(A, A[i], i)
		}
		if res == -1 {
			return 0
		}
		i = res
		jumps++
	}

	if i == len(A)-1 {
		fmt.Println("ok")
		return 1
	} else {
		return 0
	}
}

func oddEvenJumps(A []int) int {
	oddJumps = map[int]int{}
	evenJumps = map[int]int{}
	c := 0
	for i := range A {
		fmt.Println(i)
		c += oddEvenJumpsR(A, i)
	}
	return c
}

func main() {

	in := []int{10, 13, 12, 14, 15}

	r := oddJump(in, 14, 3)
	fmt.Println(r)
}
