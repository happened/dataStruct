package bTree

import "fmt"

type BinaryTreeNodeType float32

//func (src *BinaryTreeNodeType)Compare(dst *BinaryTreeNodeType) bool{
//	return true
//}

type BinaryTreeNode struct {
	val BinaryTreeNodeType
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

//二叉树 左小右大
func (bt *BinaryTree)InsertNode(val BinaryTreeNodeType)*BinaryTree{
	pre:=&BinaryTreeNode{}
	locate:=bt.root
	forward:=true

	for locate!=nil{
		if val>locate.val{
			pre=locate
			locate=locate.right
			forward=true
		}else if val==locate.val{
			fmt.Println("val exist")
			return bt
		}else{
			pre=locate
			locate=locate.left
			forward=false
		}
	}

	locate=&BinaryTreeNode{
		val: val,
		left: nil,
		right: nil,
	}
	if	forward{
		pre.right=locate
	}else {
		pre.left=locate
	}
	return bt
}

func CreateBinaryTree(rootVal BinaryTreeNodeType)*BinaryTree{
	return &BinaryTree{
		root: &BinaryTreeNode{
			val: rootVal,
			left:nil,
			right: nil,
		},
	}
}





