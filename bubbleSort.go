package fullsort

import (
	"time"
	"errors"
)

func BubbleSortInt(arr []int) (*Response, error) {

	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	} else if len(arr) == 1 {
		return &Response{0, 0, 0}, nil
	}

	var swaps, accesses uint32
	tstart := time.Now()

	var i int
	done := false
	for !done {
		done = true
		for i = 0; i < len(arr)-1; i++ {
			accesses += 2
			if arr[i] > arr[i+1] {
				swaps++
				arr[i], arr[i+1] = arr[i+1], arr[i]
				done = false
			}
		}
	}
	return &Response{swaps, accesses, time.Since(tstart)}, nil
}

func BubbleSort(arr interface{}, eval func(a, b interface{}) bool) (*Response, error) {
	return nil, nil
}