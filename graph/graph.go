package graph

import "fmt"

type Graph struct {
	edges      [][]int
	vertexList []string
	numOfEdges int
}

func NewGraph(n int) *Graph {
	var edges [][]int
	vertexList := make([]string, 0)
	graph := &Graph{
		edges:      edges,
		vertexList: vertexList,
		numOfEdges: 0,
	}
	
	return graph
}

//新增顶点
func (g *Graph)insertVertex(vertex string)  {
	g.vertexList = append(g.vertexList, vertex)
}

//新增边 v1代表第一个顶点下标  v2代表第二个顶点下标
func (g *Graph)insertEdge(v1, v2, weight int)  {
	g.edges[v1][v2] = weight
	g.edges[v2][v1] = weight
	g.numOfEdges++
}

//得到便的数目
func (g *Graph)getNumOfVertex() int {
	return g.numOfEdges
}

func (g *Graph)getValueByIndex(index int) string {
	return g.vertexList[index]
}

func (g *Graph)showGraph()  {
	for k,v := range g.edges {
		for _, v1 := range v {
			fmt.Println("k:",k, v1)
		}
	}
}