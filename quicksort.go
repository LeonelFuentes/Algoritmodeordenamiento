package main

import "fmt"

func main() {
	var array = []int{9, 7, 1, 6, 5, 4, 3, 2, 8, 0}

	fmt.Println(array)
	qsort(array)
	fmt.Println(array)
}

func qsort(array []int) {
	qsortRec(array, 0, len(array)-1)
}

//Quicksort recursive
func qsortRec(array []int, left, right int) {
	//Base case
	if left < right {
		var p = partition(array, left, right)
		qsortRec(array, left, p-1)
		qsortRec(array, p+1, right)
	}
}

//Calcula la particion y coloca los elementos <pivote a la izquierda y> pivote a la derecha
func partition(array []int, left, right int) int {
	var pivot = array[getPivot(array, left, right)]
	var l, r = left, right

	for l < r {
		for array[l] < pivot {
			l++
		}

		for array[r] > pivot {
			r--
		}

		if l < r {
			array[l], array[r] = array[r], array[l]
		}
	}

	return l //The partition
}

//Calculate the index of the pivot by selecting the element in the middle
func getPivot(array []int, left, right int) int {
	return left + ((right - left) / 2)
}
