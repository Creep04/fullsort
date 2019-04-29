package fullsort

import (
	"testing"
	"math/rand"
)

func randomIntArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i+1
	}
	rand.Shuffle(size, func(i, j int){
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func isInOrderInt(arr []int) bool {
	for i, n := range arr {
		if n != i+1 {
			return false
		}
	}
	return true
}

func TestBubbleSortInt(t *testing.T) {

	// Easy test
	arr := randomIntArray(10)
	res, err := BubbleSortInt(arr)
	if err != nil {
		t.Error("order 10 integers: returned an error:", err)
	}
	if !isInOrderInt(arr) {
		t.Error("order 10 integers: array isn't in order:", arr)
	}

	// Empty test
	res, err = BubbleSortInt([]int{})
	if res != nil {
		t.Error("order 0 integers: returned a response:", res)
	}
	if err == nil {
		t.Error("order 0 integers: didn't return a error")
	}

	// Medium test
	arr = randomIntArray(10000)
	res, err = BubbleSortInt(arr)
	if err != nil {
		t.Error("order 10,000 integers: returned an error:", err)
	}
	if !isInOrderInt(arr) {
		t.Error("order 10,000 integers: array isn't in order:", arr)
	}

	// One element test
	arr = randomIntArray(1)
	res, err = BubbleSortInt(arr)
	if err != nil {
		t.Error("order 1 integer: returned an error:", err)
	}
	if !isInOrderInt(arr) {
		t.Error("order 1 integer: array isn't in order:", arr)
	}
}