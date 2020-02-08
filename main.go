package main

import (
	selectionSort "github.com/rohitjb/datastructure/sorting/selection"
)

func main() {
	//sort.BubbleSortAlgo([]int{8, 5, 7, 3, 2})
	//insertionSort.Sort([]int{8, 5, 7, 3, 2, 1, 6, 4})
	selectionSort.Sort([]int{8, 5, 7, 3, 2, 1, 6, 4})
}