package _206_Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针的方式实现
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//递归的方式实现
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	return reverse(pre, cur)
}
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	pre = cur
	cur = tmp
	return reverse(pre, cur)
}
