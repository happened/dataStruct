package bTree

import (
	"fmt"
)

func (bt *BinaryTree)PreTraverse(){
	preTraverse(bt.root)
}

//LGR
func preTraverse(root *BinaryTreeNode){
	if root==nil {
		return
	}
	preTraverse(root.left)
	fmt.Println(root.val)
	preTraverse(root.right)
}
