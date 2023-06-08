package graph

import (
	"fmt"
)

type Graph struct {
	edges      [][]int  //图的二位坐标集合
	vertexList []string //顶点的集合
	numOfEdges int      //边的个数
	isVisited  []bool   //顶点是否被访问记录
}

func NewGraph(n int) *Graph {
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]int, n)
	}
	vertexList := make([]string, 0)
	isVisited := make([]bool, n)
	graph := &Graph{
		edges:      edges,
		vertexList: vertexList,
		numOfEdges: 0,
		isVisited:  isVisited,
	}

	return graph
}

// InsertVertex 新增顶点
func (g *Graph) InsertVertex(vertex string) {
	g.vertexList = append(g.vertexList, vertex)
}

// InsertEdge 新增边 v1代表第一个顶点下标  v2代表第二个顶点下标
func (g *Graph) InsertEdge(v1, v2, weight int) {
	g.edges[v1][v2] = weight
	g.edges[v2][v1] = weight
	g.numOfEdges++
}

// GetNumOfVertex 得到顶点的数目
func (g *Graph) GetNumOfVertex() int {
	return len(g.vertexList)
}

// GetNumOfEdges  得到边的数目
func (g *Graph) GetNumOfEdges() int {
	return g.numOfEdges
}

// GetValueByIndex 根据顶点的坐标获取其值
func (g *Graph) GetValueByIndex(index int) string {
	return g.vertexList[index]
}

func (g *Graph) showGraph() {
	for k, v := range g.edges {
		for k1, _ := range v {
			//打印矩阵
			fmt.Print(g.edges[k][k1], "  ")
		}
		fmt.Print("\n")
	}
}

//以下为深度优先遍历

//获取该节点的第一个临近节点的坐标, 未找到则返回-1
func (g *Graph) getFirstNeighbor(index int) int {
	sum := g.GetNumOfVertex()
	for j := 0; j < sum; j++ {
		if g.edges[index][j] == 1 {
			return j
		}
	}
	return -1
}

//根据前一个临结节点的下标来获取下一个临结节点的下表 v1的临结节点的下一个临结节点
func (g *Graph) getNextNeighbor(v1, v2 int) int {
	for j := v2 + 1; j < g.GetNumOfVertex(); j++ {
		if g.edges[v1][j] > 0 {
			return j
		}
	}
	return -1
}

//深度优先遍历  第一次index为0
func (g *Graph) dfs(index int) {
	fmt.Print(g.GetValueByIndex(index) + "->")
	g.isVisited[index] = true
	//找到节点的第一个节点
	vIndex := g.getFirstNeighbor(index)
	for vIndex != -1 {
		//判断该节点没有访问
		if !g.isVisited[vIndex] {
			g.dfs(vIndex)
		}

		//该节点已经访问
		vIndex = g.getNextNeighbor(index, vIndex)
	}
	return
}

// Dfs 深度优先遍历 把所有的节点全部遍历一边
func (g *Graph) Dfs() {
	for i := 0; i < g.GetNumOfVertex(); i++ {
		//此处要判断是否已经遍历,否则重复遍历
		if !g.isVisited[i] {
			g.dfs(i)
		}
	}
	fmt.Print("\n")
}

// Bfs 以下为广度优先算法
func (g *Graph) bfs(index int) {
	//表示队列的头节点对应下标
	var u int
	//临结节点的下标
	var w int
	//双向列表来代表队列  记录节点访问顺序
	//list := list.New()
	fmt.Print(g.GetValueByIndex(index) + "->")
	//标记已访问
	g.isVisited[index] = true
	//放置队列最前面
	//list.PushBack(index)
	//队列不能为空
	//for list.Len() > 0 {
	//取出最前面的数据,remove移除某个元素并返回元素的值
	//	u = list.Remove(list.Front()).(int)
	//得到第一个邻结节点的下标
	w = g.getFirstNeighbor(u)
	for w != -1 {
		//判断该节点是否访问过
		if !g.isVisited[w] {
			fmt.Print(g.GetValueByIndex(w) + "->")
			g.isVisited[w] = true
			//入队 放到最后面
			//			list.PushBack(w)
		}
		//找到u行 w节点的下一个节点
		w = g.getNextNeighbor(u, w)
	}
	//}
}

// Bfs 以下为广度优先算法 遍历所有节点
func (g *Graph) Bfs() {
	for i := 0; i < g.GetNumOfVertex(); i++ {
		//此处要判断是否已经遍历,否则重复遍历
		if !g.isVisited[i] {
			g.bfs(i)
		}
	}
	fmt.Print("\n")
}
