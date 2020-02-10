package quick

func Sort(arr []int, l int, h int) {
	if l < h {
		j:= partition(arr, l, h)
		Sort(arr, l, j)
		Sort(arr, j+1, h)
	}
}

func partition(arr []int, l int, h int) int {
	pivot := arr[l]
	// look for larger element than pivot
	i := l
	j := h
	for{
		for {
			i = i + 1
			if i >= len(arr) || arr[i] >= pivot {
				break
			}
		}
		// look for smaller element than pivot
		for {
			j = j - 1
			if j <= 0 || arr[j] <= pivot {
				break
			}
		}
		// swap with i and j
		if i < j  {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}

		if i > j {
			break
		}
	}

	temp := arr[l]
	arr[l] = arr[j]
	arr[j] = temp

	return j
}