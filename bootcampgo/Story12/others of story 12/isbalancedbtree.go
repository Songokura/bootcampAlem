package bootcamp

import (
	"bootcamp/btree"
)

func IsBalancedBtree(b *btree.BTree) bool {
	return isBalanced(b.Root)
}

func isBalanced(node *btree.BTreeNode) bool {
	if node == nil {
		return true
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}

func getHeight(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
// 	tree.ReplaceOrInsert(60)
// 	tree.ReplaceOrInsert(80)
// 	/* Btree Visualization:
// 	          50
// 	        /    \
// 	      30      70
// 	     /       /  \
// 	   20       60   80
// 	*/

// 	fmt.Println(IsBalancedBtree(tree)) // true
// }
