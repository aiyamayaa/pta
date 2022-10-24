package _024_Swap_Nodes_in_Pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next.Next
		cur.Next.Next.Next = cur.Next
		cur.Next = cur.Next.Next
		cur.Next.Next.Next = tmp
		cur = cur.Next.Next
	}

	return dummyHead.Next
}
