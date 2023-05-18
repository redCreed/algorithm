package graph

import "testing"

func TestNewGraph2(t *testing.T) {
	g := NewGraph(8)
	g.vertexList = append(g.vertexList, "1", "2", "3", "4", "5", "6", "7", "8")
	//添加边  A-B A-C B-C B-D B-E
	g.InsertEdge(0, 1, 1)
	g.InsertEdge(0, 2, 1)
	g.InsertEdge(1, 3, 1)
	g.InsertEdge(1, 4, 1)
	g.InsertEdge(3, 7, 1)
	g.InsertEdge(4, 7, 1)
	g.InsertEdge(2, 5, 1)
	g.InsertEdge(2, 6, 1)
	g.InsertEdge(5, 6, 1)
	g.showGraph()
	//深度优先算法
	//g.Dfs()
	//广度
	g.Bfs()
}
