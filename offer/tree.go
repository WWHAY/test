package offer

import "fmt"

//TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root.Left != nil {
		result = append(result, root.Left.Val)
	}
	if root.Right != nil {
		result = append(result, root.Right.Val)
	}
	return result
}

func get(root *TreeNode)
