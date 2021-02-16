package sort

//QuickSort .
func QuickSort(intList []int) {
	// 如果长度小于等于1就直接结束
	if len(intList) <= 1 {
		return
	}
	// 1. 将第一个值选定为基准值
	flag := intList[0]
	// 定义左右指针
	left, right := 0, len(intList)-1

	for i := 1; i <= right; {
		if intList[i] > flag {
			intList[i], intList[right] = intList[right], intList[i]
			right--
		} else {
			intList[i], intList[left] = intList[left], intList[i]
			i++
			left++
		}
	}
	// 递归
	QuickSort(intList[:left])
	QuickSort(intList[left+1:])

}

//QuickSort1 .
func QuickSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return arr
}

//BubbleSort .
func BubbleSort(arr []int) {

}
