package linkList

import (
	"fmt"
)

type doubleLinkNode struct {
	value    int
	preNode  *doubleLinkNode
	nextNode *doubleLinkNode
}

func (headNode *doubleLinkNode) create(array []int) {

	curNode := headNode

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生错误:", r)

		}
	}()

	for _, v := range array {
		tempNode := new(doubleLinkNode)
		tempNode.value = v

		curNode.nextNode = tempNode

		tempNode.preNode = curNode
		tempNode.nextNode = nil

		curNode = tempNode
	}
}

func (headNode *doubleLinkNode) display() {
	if headNode == nil || headNode.nextNode == nil {
		fmt.Println("此双链表为空")
		return
	}

	node := headNode.nextNode

	for ; node != nil; node = node.nextNode {
		fmt.Println(node.value)
	}

}

func (headNode *doubleLinkNode) isExist(value int) bool {
	return false
}

func (headNode *doubleLinkNode) findValue(value int) interface{} {
	if headNode == nil || headNode.nextNode == nil {
		return "无法从一个空链表中查找数据"
	}

	node := headNode.nextNode

	for ; node != nil; node = node.nextNode {
		if node.value == value {
			return node
		}
	}
	return "未找到该元素值"
}

func TestDoubleLinkList() {
	array := []int{1, 5, 4, 6, 7}

	tempnode := new(doubleLinkNode)
	tempnode.create(array)

	var allNode node = tempnode

	//allNode.display()

	pointNode := allNode.findValue(11)

	switch result := pointNode.(type) {
	case string:
		fmt.Println(result)
	case *doubleLinkNode:
		fmt.Println(result.preNode.value)
	}

}
