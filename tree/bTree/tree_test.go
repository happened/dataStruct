package bTree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree:=CreateBinaryTree(10).
		InsertNode(3234).InsertNode(1.42143).InsertNode(20).InsertNode(3).InsertNode(8).
		InsertNode(-0.12312)

	//前序遍历
	tree.PreTraverse()

	//判断是否存在某个值
	fmt.Println(tree.Exist(3234))
}
