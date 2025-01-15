package btree

func (b *BTree) Get(value int) *BTreeNode {
	current := b.Root

	for current != nil {
		if value == current.Value {
			return current
		} else if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}
