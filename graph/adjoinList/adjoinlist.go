package adjoinList

// 邻接表 有向无权图

type nodeType=int

type node struct {
	val nodeType
	next *node
}

type Graph struct {
	nodes []*node
}

func CreateGraph() *Graph {
	return &Graph{}
}

func (g *Graph)addEdge(src nodeType,dst nodeType)*Graph {
	for _,tempPoint:=range g.nodes{
		if tempPoint.val==src{
			curNode:=tempPoint
			for curNode.next!=nil{
				if curNode.val==dst{
					return g
				}
				curNode=curNode.next
			}
			curNode.next=&node{
				val:dst,
				next:nil,
			}
			return g
		}
	}

	g.nodes=append(g.nodes,&node{
		val: src,
		next: &node{
			val: dst,
			next: nil,
		},
	})
	return g
}
