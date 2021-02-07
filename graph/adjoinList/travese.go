package adjoinList

import "fmt"

func TraverseAdJoinListGraph(g *Graph){
	for _,node:=range g.nodes{
		for node!=nil{
			fmt.Printf("%v --> ",node.val)
			node=node.next
		}
		fmt.Println("------------")
	}
}
