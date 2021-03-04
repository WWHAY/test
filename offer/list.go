package offer

//ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var newList *ListNode
	for head != nil {
		node := head.Next
		head.Next = newList
		newList = head
		head = node
	}
	var result []int
	for newList != nil {
		result = append(result, newList.Val)
		newList = newList.Next
	}
	return result
}

func reverse1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}
	return []int{head.Val}
}
