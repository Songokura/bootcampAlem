package bootcamp

import (
	"bootcamp/btree"
)

func GetMin(b *btree.BTree) *btree.BTreeNode {
	current := b.Root

	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

// func main() {
// 	tree := btree.NewBTree()
// 	tree.ReplaceOrInsert(50)
// 	tree.ReplaceOrInsert(20)
// 	tree.ReplaceOrInsert(40)
// 	tree.ReplaceOrInsert(30)
// 	tree.ReplaceOrInsert(70)
// 	/* Btree Visualization:
// 	       50
// 	     /    \
// 	   20     70
// 	     \
// 	      40
// 	     /
// 	   30
// 	*/

// 	minNode := GetMin(tree)
// 	if minNode != nil {
// 		fmt.Println(minNode.Value) // 20
// 	}
// }
