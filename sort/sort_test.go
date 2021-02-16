package sort

import "testing"

func TestQuick(t *testing.T) {
	var arr = []int{2, 6, 9, 5, 4, 1, 3}
	QuickSort(arr)
	t.Log(arr)
}
