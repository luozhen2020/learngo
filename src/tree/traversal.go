package tree

func (node *Node) Traversal() {
	if node == nil {
		return
	}

	node.Left.Traversal()
	node.Print()
	node.Right.Traversal()
}
