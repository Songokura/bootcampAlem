package bootcamp

import (
	"bootcamp/btree"
)

func GetMax(b *btree.BTree) *btree.BTreeNode {
	current := b.Root

	for current != nil && current.Right != nil {
		current = current.Right
	}
	return current
}

// func main() {
// 	tree := btree.NewBTree()
// 	tree.ReplaceOrInsert(50)
// 	tree.ReplaceOrInsert(20)
// 	tree.ReplaceOrInsert(40)
// 	tree.ReplaceOrInsert(30)
// 	tree.ReplaceOrInsert(10)
// 	tree.ReplaceOrInsert(90)
// 	/* Btree Visualization:
// 	          50
// 	        /    \
// 	      20      90
// 	     /  \
// 	   10    40
// 	        /
// 	      30
// 	*/

// 	maxNode := GetMax(tree)
// 	if maxNode != nil {
// 		fmt.Println(maxNode.Value) // 90
// 	}
// }
