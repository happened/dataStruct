package bTree

func (bt *BinaryTree) Exist(val BinaryTreeNodeType)(bool){
	node:=bt.root
	for node!=nil{
		if node.val==val{
			return true
		}else if node.val >val{
			node=node.left
		}else{
			node=node.right
		}
	}
	return false
}