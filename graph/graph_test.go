package graph

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	g.vertexList = append(g.vertexList, "A","B","C","D","E")
	g.insertEdge(0,1,1)
	g.insertEdge(0,2,1)
	g.insertEdge(1,2,1)
	g.insertEdge(1,3,1)
	g.insertEdge(1,4,1)
	g.showGraph()
}
