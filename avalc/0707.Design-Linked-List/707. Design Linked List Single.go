package _707_Design_Linked_List

type MyLinkedList struct {
	Head *ListNode
	Size int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}

}

//index为0的时候，size为1
func (this *MyLinkedList) Get(index int) int {
	cur := this.Head
	if index < 0 || index >= this.Size {
		return -1
	}
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	index = max(index, 0)
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode := ListNode{Val: val, Next: cur.Next}
	cur.Next = &newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
