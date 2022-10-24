package _019_Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre, cur := dummyHead, dummyHead
	for i := 0; i <= n; i++ {
		if cur == nil {
			return dummyHead.Next
		}
		cur = cur.Next
	}
	for ; cur != nil; cur = cur.Next {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
