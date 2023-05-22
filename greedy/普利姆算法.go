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
