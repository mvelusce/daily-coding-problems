package main

import "math/rand"

func QuickSort(arr []int) []int {

	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}
	p, a := partition(arr, lo, hi)
	a = quickSort(a, lo, p-1)
	return quickSort(a, p+1, hi)
}

func partition(arr []int, lo int, hi int) (int, []int) {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
		}
	}
	tmp := arr[hi]
	arr[hi] = arr[i]
	arr[i] = tmp
	return i, arr
}

func main() {

	arr := [100]int{}
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
		print(arr[i], " ")
	}
	println()
	res := QuickSort(arr[:])
	for _, el := range res {
		print(el, " ")
	}
}
