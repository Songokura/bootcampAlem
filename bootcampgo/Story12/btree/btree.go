package btree

type BTree struct {
	Root *BTreeNode
}

type BTreeNode struct {
	Parent      *BTreeNode
	Left, Right *BTreeNode
	Value       int
}

func NewBTree() *BTree {
	return &BTree{}
}

func (b *BTree) ReplaceOrInsert(v int) {
	if b.Root == nil {
		b.Root = &BTreeNode{Value: v}
	} else {
		insertNode(b.Root, v)
	}
}

func insertNode(node *BTreeNode, v int) {
	if v < node.Value {
		if node.Left == nil {
			node.Left = &BTreeNode{Parent: node, Value: v}
		} else {
			insertNode(node.Left, v)
		}
	} else if v > node.Value {
		if node.Right == nil {
			node.Right = &BTreeNode{Parent: node, Value: v}
		} else {
			insertNode(node.Right, v)
		}
	}
}
