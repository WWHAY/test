package tree

import "test/lib"

func preorderTraversal(root *lib.TreeNode) (res []int) {
	stack := []*lib.TreeNode{}
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
