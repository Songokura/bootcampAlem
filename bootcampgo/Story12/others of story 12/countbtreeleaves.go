package bootcamp

import (
	"bootcamp/btree"
)

func CountBtreeLeaves(b *btree.BTree) int {
	if b == nil {
		return 0
	}
	return CountBtreeL(b.Root)
}

func CountBtreeL(b *btree.BTreeNode) int {
	if b == nil {
		return 0
	}
	if b.Left == nil && b.Right == nil {
		return 1 // node is a leaf node
	}
	return CountBtreeL(b.Left) + CountBtreeL(b.Right)
}

// func main() {
// 	tree := btree.NewBTree()
// 	tree.ReplaceOrInsert(50)
// 	tree.ReplaceOrInsert(30)
// 	tree.ReplaceOrInsert(70)
// 	tree.ReplaceOrInsert(20)
// 	tree.ReplaceOrInsert(40)
// 	tree.ReplaceOrInsert(60)
// 	tree.ReplaceOrInsert(10)
// 	/* Btree Visualization:
// 	             50
// 	           /    \
// 	         30      70
// 	        /  \    /
// 	      20   40  60
// 	     /
// 	   10
// 	*/

// 	// Number of leaf nodes in the tree:
// 	leafCount := CountBtreeLeaves(tree)
// 	fmt.Println(leafCount) // 3
// }
