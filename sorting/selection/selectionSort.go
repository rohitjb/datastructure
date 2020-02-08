package selection

import "fmt"

func Sort(arr []int) {
	for index, element := range arr {
		k:= index
		min:= element

		for j:= index+1 ; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				k = j
			}
		}

		temp := arr[index]
		arr[index] = arr[k]
		arr[k] = temp
	}
	fmt.Println(arr)
}