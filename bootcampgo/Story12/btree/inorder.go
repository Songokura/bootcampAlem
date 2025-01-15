package btree

func (b *BTree) InOrderTraversal(fn func(n *BTreeNode)) {
	inOrder(b.Root, fn)
}

func inOrder(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}

	inOrder(node.Left, fn)
	fn(node)
	inOrder(node.Right, fn)
}
