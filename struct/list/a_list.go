package list

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) InsertTail(node *Node) {
	n.Next = node
}

func (n *Node) InsertHead(node *Node) *Node {
	node.Next = n
	return node
}

func (n *Node) Insert(node *Node, insert int) {
	m := n
	for i := 0; i < insert; i++ {
		m = m.Next
	}
	f := m.Next
	m.Next = node
	node.Next = f
}
func Delete() {

}
