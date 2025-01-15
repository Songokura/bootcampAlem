package bootcamp

import (
	"bootcamp/btree"
)

// CountBtreeNodes counts the total number of nodes in the BST
func CountBtreeExtend(b *btree.BTreeNode) int {
	if b == nil {
		return 0
	}

	return 1 + CountBtreeExtend(b.Left) + CountBtreeExtend(b.Right)
}

func CountBtreeNodes(b *btree.BTree) int {
	if b == nil {
		return 0
	}
	return CountBtreeExtend(b.Root)
}

// func main() {
// 	tree := btree.NewBTree()
// 	tree.ReplaceOrInsert(50)
// 	tree.ReplaceOrInsert(30)
// 	tree.ReplaceOrInsert(70)
// 	tree.ReplaceOrInsert(20)
// 	tree.ReplaceOrInsert(40)
// 	/* Btree Visualization:
// 	          50
// 	        /    \
// 	      30     70
// 	     /  \
// 	   20    40
// 	*/

// 	fmt.Println("Total number of nodes in the tree:")
// 	totalNodes := CountBtreeNodes(tree)
// 	fmt.Println(totalNodes) // 5
// }
