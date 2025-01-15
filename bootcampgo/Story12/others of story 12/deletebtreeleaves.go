package bootcamp

import (
	"bootcamp/btree"
)

func DeleteBtreeLeaves(b *btree.BTree) {
	if b.Root == nil {
		return
	}
	deleteLeaves(b.Root)
}

func deleteLeaves(node *btree.BTreeNode) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		parent := node.Parent
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			} else if parent.Right == node {
				parent.Right = nil
			}
		}
		return
	}
	deleteLeaves(node.Left)
	deleteLeaves(node.Right)
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
// 	         50
// 	       /    \
// 	     30      70
// 	    /  \    /  \
// 	   20  40  60  80
// 	*/

// 	// Tree before deleting leaves:
// 	printTree(tree.Root) // 20 30 40 50 60 70 80

// 	DeleteBtreeLeaves(tree)
// 	/* Btree Visualization:
// 	       50
// 	     /    \
// 	   30      70
// 	*/

// 	// Tree after deleting leaves:
// 	printTree(tree.Root) // 30 50 70
// }

// func printTree(node *btree.BTreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	printTree(node.Left)
// 	fmt.Println(node.Value)
// 	printTree(node.Right)
// }
