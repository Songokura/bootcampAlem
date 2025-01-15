package bootcamp

import (
	"bootcamp/btree"
)

func LevelsBtree(b *btree.BTree) int {
	if b.Root == nil {
		return 0
	}
	return calculateLevels(b.Root)
}

func calculateLevels(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}
	leftLevels := calculateLevels(node.Left)
	rightLevels := calculateLevels(node.Right)

	return 1 + max(leftLevels, rightLevels)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	tree := btree.NewBTree()
// 	tree.ReplaceOrInsert(50)
// 	tree.ReplaceOrInsert(30)
// 	tree.ReplaceOrInsert(70)
// 	tree.ReplaceOrInsert(20)
// 	tree.ReplaceOrInsert(40)
// 	tree.ReplaceOrInsert(60)
// 	tree.ReplaceOrInsert(80)
// 	/* Btree Visualization:
// 	           50          level 0
// 	         /    \
// 	       30      70      level 1
// 	      /  \    /  \
// 	     20  40  60  80    level 2

// 	   total number of levels: 3
// 	*/

// 	numLevels := LevelsBtree(tree)
// 	fmt.Println(numLevels) // 3
// }
