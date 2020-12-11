package main

import "math/rand"

func MergeSort(arr []int) []int {

	return arr
}

func main() {

	arr := [100]int{}
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
		print(arr[i], " ")
	}
	println()
	res := MergeSort(arr[:])
	for _, el := range res {
		print(el, " ")
	}
}
