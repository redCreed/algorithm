package graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	g.vertexList = append(g.vertexList, "A", "B", "C", "D", "E")
	//添加边  A-B A-C B-C B-D B-E
	g.InsertEdge(0, 1, 1)
	g.InsertEdge(0, 2, 1)
	g.InsertEdge(1, 2, 1)
	g.InsertEdge(1, 3, 1)
	g.InsertEdge(1, 4, 1)
	g.showGraph()
	//深度优先算法
	g.Dfs()
	g.Bfs()
}
