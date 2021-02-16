package sort

import "testing"

func TestQuick(t *testing.T) {
	var arr = []int{2, 6, 9, 5, 5, 4, 1, 4, 3}
	QuickSort1(arr)
	t.Log(arr)
}
