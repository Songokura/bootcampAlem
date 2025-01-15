package btree

func (b *BTree) PostOrderTraversal(fn func(n *BTreeNode)) {
	postOrder(b.Root, fn)
}

func postOrder(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}
	postOrder(node.Left, fn)
	postOrder(node.Right, fn)
	fn(node)
}
