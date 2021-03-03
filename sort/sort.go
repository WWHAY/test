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
func QuickSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}
	flag := arr[0]
	l, r := 0, len(arr)-1
	for l < r {
		for arr[r] > flag {
			r--
		}
		if arr[r] < flag {
			arr[r], arr[l] = arr[l], arr[r]
		}
		for arr[l] < flag {
			l++
		}
		if arr[l] > flag {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	QuickSort(arr[:l])
	QuickSort(arr[l+1:])
}

//BubbleSort .
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
