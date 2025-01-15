package btree

func (b *BTree) ApplyByLevel(fn func(node *BTreeNode, level int)) {
	if b.Root == nil {
		return
	}

	currentLevel := []*BTreeNode{b.Root}
	level := 0

	for len(currentLevel) > 0 {
		nextLevel := []*BTreeNode{}

		for _, node := range currentLevel {
			fn(node, level)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		currentLevel = nextLevel
		level++
	}
}
