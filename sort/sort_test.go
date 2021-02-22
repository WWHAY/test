package sort

import "testing"

func TestQuick(t *testing.T) {
	var arr = []int{2, 6, 9, 5, 5, 4, 1, 4, 3}
	QuickSort(arr)
	t.Log(arr)
}

func TestBubble(t *testing.T) {
	var arr = []int{2, 6, 9, 5, 5, 4, 1, 4, 3}
	BubbleSort(arr)
	t.Log(arr)
}
