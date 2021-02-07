package adjoinList

import "testing"


func TestTraverseAdJoinListGraph(t *testing.T) {

	g:=CreateGraph()
	g.addEdge(1,2)
	g.addEdge(2,3)
	g.addEdge(2,4)

	g.addEdge(19,4)
	TraverseAdJoinListGraph(g)
}