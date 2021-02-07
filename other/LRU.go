package other

import (
	"container/list"
	"sync"
)

// 实现golang版本的LRU算法

type node struct {
	key,value string
}


type Cache struct {
	maxLength int
	list *list.List
	lock *sync.Mutex

	cache map[string]*list.Element

}

func NewCache(length int) *Cache{
	return &Cache{
		maxLength:length,
		list:list.New(),
		cache:make(map[string]*list.Element),
	}
}


func (c *Cache) Add(nodekey,nodeValue string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if element, exist := c.cache[nodekey]; exist {
		c.list.MoveToFront(element)
		element.Value.(*node).value = nodeValue
		return
	}
	re := c.list.PushFront(node{key: nodekey, value: nodeValue})
	c.cache[nodekey] = re
	if c.maxLength<c.list.Len(){
		last:=c.list.Back()
		c.list.Remove(last)
	}
}

