package greedy

import "fmt"

/*
	普利姆算法：求最小生成树,也就是在包含n个顶点的连通图中,找出只有(n-1)条边包含所有n个顶点的连通子图,也就是所谓的极小连通子图
*/

type MinTree struct {
}

//创建最小的图路径
func (m *MinTree) createGraph(mg *MGraph, data []string, weight [][]int) {
	for i := 0; i < len(data); i++ {
		mg.vertex[i] = data[i]

		for j := 0; j < len(data); j++ {
			mg.weight[i][j] = weight[i][j]
		}
	}
	for k, v := range mg.weight {
		for k1, _ := range v {
			//打印矩阵
			fmt.Print(mg.weight[k][k1], "  ")
		}
		fmt.Print("\n")
	}

	m.prim(mg, 0)
}

func (m *MinTree) prim(mg *MGraph, v int) {
	//判断当前节点是否已经访问
	isVisited := make(map[int]bool)
	//标记当前节点已访问
	isVisited[v] = true
	//h1和h2记录2个顶点下标
	h1 := -1
	h2 := -1
	//默认最小权重为10000
	minWeight := 10000

	//遍历边的个数 普利姆算法是n个点，就会有n-1个边
	for k := 1; k < mg.nums; k++ {

		//每个子图和哪个节点最近
		for i := 0; i < mg.nums; i++ { //表示已经访问的节点
			for j := 0; j < mg.nums; j++ { //表示未访问的节点
				//当二维矩阵的值为10000，则表示2个节点不联通
				if isVisited[i] && !isVisited[j] && mg.weight[i][j] < minWeight {
					//记录已经访问的节点i与所有和它相邻的节点的最小权重
					minWeight = mg.weight[i][j]
					//表示已经访问的节点下标
					h1 = i
					//相邻节点的下标(未访问的)
					h2 = j
				}
			}
		}
		//将未访问的标记为已经访问
		isVisited[h2] = true
		fmt.Println(mg.vertex[h1], ",", mg.vertex[h2], ">>>", minWeight)
		//重置最小权重
		minWeight = 10000
	}
}

type MGraph struct {
	weight [][]int  //代表边
	nums   int      //顶点个数
	vertex []string //顶点集合
}

func NewMGraph(num int) *MGraph {
	mg := new(MGraph)
	mg.nums = num
	mg.vertex = make([]string, num)
	mg.weight = make([][]int, num)
	for i := 0; i < num; i++ {
		mg.weight[i] = make([]int, num)
	}
	return mg
}
