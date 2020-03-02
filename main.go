package main

import (
	"fmt"
	quickSort "github.com/rohitjb/datastructure/sorting/quick"
)

func main() {
	//sort.BubbleSortAlgo([]int{8, 5, 7, 3, 2})
	//insertionSort.Sort([]int{8, 5, 7, 3, 2, 1, 6, 4})
	//selectionSort.Sort([]int{8, 5, 7, 3, 2, 1, 6, 4})
	arr := []int{50, 70, 60, 90, 40, 80, 10, 20, 30}
	//arr := []int{2, 9, 5, 3, 6}
	quickSort.Sort(arr, 0, len(arr))
	fmt.Println(arr)
}