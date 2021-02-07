// 实现带头结点的单链表
package linkList

import (
	"fmt"
	"os"
)

type linkNode struct {
	value int
	next  *linkNode
}

//利用数组初始化链表
func (headNode *linkNode) create(num []int) {
	curNode := headNode
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生错误", r)
			os.Exit(1)
		}
	}()

	for _, number := range num {
		tempNode := new(linkNode)
		tempNode.value = number
		tempNode.next = nil
		curNode.next = tempNode
		curNode = tempNode
	}
}

//---实现node接口---//

//打印链表
func (headNode *linkNode) display() {
	if headNode == nil || headNode.next == nil {
		fmt.Println("该单链表为空")
		return
	}
	node := headNode.next
	for ; node != nil; node = node.next {
		fmt.Println(node.value)
	}
}

//判断某个值的节点是否存在
func (headNode *linkNode) isExist(value int) bool {
	if headNode == nil {
		return false
	}
	tempNode := headNode.next
	for ; tempNode != nil; tempNode = tempNode.next {
		if tempNode.value == value {
			return true
		}
	}
	return false

}

func (headNode *linkNode) findValue(value int) interface{} {
	if headNode == nil || headNode.next == nil {
		return "无法从空的单链表中查询数据"
	}
	node := headNode.next
	for ; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}
	return "未找到该元素值"
}

//翻转链表
func (headNode *linkNode)reverseLinkList(){
	 defer func() {
	 	if r:=recover();r!=nil{
	 		fmt.Println(r)
		}
	 }()
	if headNode==nil || headNode.next==nil{
		fmt.Println("空链表")
	}else if headNode.next.next==nil{
		fmt.Println("只有一个元素 翻转成功")
	}else{
		p:=headNode.next
		q:=p.next
		p.next=nil

		for q!=nil{
			temp:=q.next
			q.next=p
			headNode.next=q
			p=q
			q=temp
		}

		fmt.Println("翻转成功")
	}


}



func TestSingleLinkList() {

	headNode := new(linkNode)
	array := []int{1,2,4,56}
	headNode.create(array)
	/*fmt.Println(headNode.isExist(1))
	var allNode node = headNode
	allNode.display()

	pointNode := allNode.findValue(1000)
	switch result := pointNode.(type) {
	case string:
		fmt.Println(result)
	case *linkNode:
		fmt.Println(result.value)
	}*/

	headNode.reverseLinkList()
	headNode.display()
}
