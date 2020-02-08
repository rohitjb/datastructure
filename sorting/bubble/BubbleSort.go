package bubble

import "fmt"

func BubbleSortAlgo(arr []int) {
	for index, _ := range arr {
		for j := 0; j < len(arr) - 1 - index; j++ {
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}