package offer

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	root := postorder[len(postorder)-1]
	var r int
	for i := len(postorder) - 2; i >= 0; i-- {
		if postorder[i] <= root {
			r = i
			fmt.Println("r:", i)
			break
		}
	}
	var l int
	for ; l < r; l++ {
		if postorder[l] > root {
			return false
		}
	}
	if l == r-1 {
		return true
	}
	return verifyPostorder(postorder[:r]) && verifyPostorder(postorder[r:len(postorder)-1])
}
