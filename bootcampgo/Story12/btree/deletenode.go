package btree

func (b *BTree) DeleteNode(value int) {
	b.Root = deleteNode(b.Root, value)
}

func deleteNode(node *BTreeNode, value int) *BTreeNode {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		min := node.Right
		for min.Left != nil {
			min = min.Left
		}
		node.Value = min.Value

		node.Right = deleteNode(node.Right, min.Value)
	}
	return node
}
