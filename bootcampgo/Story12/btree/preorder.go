package btree

func (b *BTree) PreOrderTraversal(fn func(n *BTreeNode)) {
	preOrder(b.Root, fn)
}

func preOrder(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}
	fn(node)
	preOrder(node.Left, fn)
	preOrder(node.Right, fn)
}
