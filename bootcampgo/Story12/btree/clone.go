package btree

func (b *BTree) Clone() *BTree {
	if b.Root == nil {
		return &BTree{}
	}

	return &BTree{Root: cloneNode(b.Root)}
}

func cloneNode(node *BTreeNode) *BTreeNode {
	if node == nil {
		return nil
	}

	clone := &BTreeNode{
		Value: node.Value,
	}

	clone.Left = cloneNode(node.Left)
	clone.Right = cloneNode(node.Right)
	if clone.Left != nil {
		clone.Left.Parent = clone
	}
	if clone.Right != nil {
		clone.Right.Parent = clone
	}

	return clone
}
