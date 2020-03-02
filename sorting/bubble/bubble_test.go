package bubble

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type expectedResult struct {
	result	[]int
}

var BubbleSortTests = []struct{
	name 		string
	input		[]int
	want 		expectedResult
}{
	{
		name: 	"SortTheUnsortedArray",
		input:	[]int{2, 9, 5, 3, 6},
		want:	expectedResult{result:[]int{2, 3, 5, 6, 9}},
	}, {
		name: 	"SortArrayOfLargeNumber",
		input:	[]int{190, 2000, 150, 2100, 500, 300, 600, 250},
		want:   expectedResult{result:[]int{150, 190, 250, 300, 500, 600, 2000, 2100},},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range BubbleSortTests {
		tt:= tt
		t.Run(tt.name, func(t *testing.T) {
			BubbleSortAlgo(tt.input)
			assert.Equal(t, tt.want.result, tt.input)
		})
	}
}