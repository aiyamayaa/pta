package _707_Design_Linked_List_double

//type MyLinkedList struct {
//	Head,Tail *Node
//	Size int
//}
//
//type Node struct{
//	Prev,Next *Node
//	Val int
//}
//
//func Constructor() MyLinkedList {
//	head := &Node{}
//	tail := &Node{}
//	head.Next = tail
//	tail.Prev = head
//	fmt.Println(fmt.Sprintf("New---Head:%+v",head))
//	fmt.Println(fmt.Sprintf("New---Tail:%+v",tail))
//	return MyLinkedList{head,tail,0}
//
//}
//
////index为0的时候，size为1
//func (this *MyLinkedList) Get(index int) int {
//	if index<0 || index >= this.Size {
//		return -1
//	}
//	var cur *Node
//	if index +1 < this.Size - index{
//		cur = this.Head
//		for i:=0;i<=index;i++ {
//			cur = cur.Next
//		}
//	}else{
//		cur = this.Tail
//		for i:=this.Size;i>index;i-- {
//			cur = cur.Prev
//		}
//	}
//	return cur.Val
//}
//
//
//func (this *MyLinkedList) AddAtHead(val int)  {
//	this.AddAtIndex(0,val)
//}
//
//
//func (this *MyLinkedList) AddAtTail(val int)  {
//	this.AddAtIndex(this.Size,val)
//}
//
//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}
//func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//	if index > this.Size {
//		return
//	}
//	var cur *Node
//	index = max(index, 0)
//	if index < this.Size - index {
//		fmt.Println(fmt.Sprintf("===from head "))
//		cur= this.Head
//		for i:=0;i<index;i++{
//			cur = cur.Next
//		}
//		fmt.Println(fmt.Sprintf("===from head==cur:%+v ",cur))
//	}else {
//		cur= this.Tail
//		fmt.Println(fmt.Sprintf("===from tail "))
//		for i:=this.Size;i>=index;i--{
//			fmt.Println(fmt.Sprintf("===from tail --"))
//			cur = cur.Prev
//
//		}
//		fmt.Println(fmt.Sprintf("===from tail==cur:%+v ",cur))
//	}
//	newNode := &Node{Val:val,Next:cur.Next,Prev:cur}
//	fmt.Println(fmt.Sprintf("%+v",cur))
//	cur.Next.Prev = newNode
//	cur.Next = newNode
//	this.Size++
//
//}
//
//
//func (this *MyLinkedList) DeleteAtIndex(index int)  {
//	if index >= this.Size||index<0 {
//		return
//	}
//	var cur *Node
//
//	if index+1>this.Size-index{
//		cur= this.Head
//		for i:=0;i>this.Size;i++{
//			cur = cur.Next
//		}
//	}else{
//		cur= this.Tail
//		for i:=this.Size;i>=index;i--{
//			cur = cur.Prev
//		}
//	}
//	cur.Next.Next.Prev = cur
//	cur.Next = cur.Next.Next
//	this.Size--
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
