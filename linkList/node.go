//这个包主要实现链表结构 如 循环链表 单链表 双链表等
package linkList

//结点接口 保留一些都会用到的方法

type node interface {
	//create(arr []int)
	display()	//打印链表内容
	isExist(value int)bool//判断某个值是否存在
	findValue(value int )interface{}	//返回特定的值所在的节点
}

