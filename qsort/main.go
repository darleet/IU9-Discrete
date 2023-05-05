package main

import "fmt"

var testArr1 = [10]int{1, 9, 2, 8, 3, 7, 5, 6, 0, 1}
//var testArr1 = [0]int{}

func less(i, j int) bool {
	return int64(testArr1[i]) < int64(testArr1[j])
}

func swap(i, j int) {
	testArr1[i], testArr1[j] = testArr1[j], testArr1[i]
}

func partition(left, right int,
	less func(i, j int) bool,
	swap func(i, j int)) int {
	i := left
	for j := left; j < right; j++ {
		if less(j, right) {
			swap(i, j)
			i++
		}
	}
	swap(i, right)
	return i
}

func qsortrec(left, right int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	if left < right {
		median := partition(left, right, less, swap)
		qsortrec(left, median-1, less, swap)
		qsortrec(median+1, right, less, swap)
	}
}

func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	qsortrec(0, n-1, less, swap)
}

func main() {
	qsort(10, less, swap)
	//qsort(0, less, swap)
	fmt.Print(testArr1)
}
