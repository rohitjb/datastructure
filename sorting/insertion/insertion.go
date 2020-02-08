package insertion

import "fmt"

func Sort(arr []int) {
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		j := i - 1
		element := arr[i]
		for j >= 0 && arr[j] > element {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = element
	}
	fmt.Println(arr)
}
