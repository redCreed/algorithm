package greedy

import "fmt"

/*
	弗洛伊德算法: 计算每个顶点到其他顶点的最短路径(所有顶点而非指定顶点)
	迪杰斯特拉算法: 计算指定顶点到其他顶点的最短路径
*/

type FGraph struct {
	vertex []string //存放顶点数
	dis    [][]int  //存放各个顶点到其他顶点的距离
	pre    [][]int  //到达目标顶点的前驱顶点
}

func NewFGraph(length int, data [][]int, vertex []string) *FGraph {
	fg := new(FGraph)
	fg.vertex = vertex
	fg.dis = data
	fg.pre = make([][]int, length)
	for i := 0; i < length; i++ {
		dd := make([]int, length)
		for j := 0; j < length; j++ {
			dd[j] = i
		}
		fg.pre[i] = dd
	}

	return fg
}

func (fg *FGraph) print() {
	length := len(fg.vertex)
	for k := 0; k < length; k++ {
		for j := 0; j < length; j++ {
			fmt.Print(fg.vertex[fg.pre[k][j]], "  ")
		}
		fmt.Println("")
		for i := 0; i < length; i++ {
			fmt.Print(fg.dis[k][i], "  ")
		}
		fmt.Println("")
	}
}

func (fg *FGraph) Floyd() {
	//临时保存2个点之前的距离
	tempL := 0
	length := len(fg.vertex)
	//中间顶点遍历
	for k := 0; k < length; k++ {
		//从i顶点开始出发
		for i := 0; i < length; i++ {
			//到达j顶点
			for j := 0; j < length; j++ {
				tempL = fg.dis[i][k] + fg.dis[k][j] //从i顶点出发，经过k中间顶点，最后到达j顶点
				if tempL < fg.dis[i][j] {           //求i到j的最小值
					fg.dis[i][j] = tempL
					fg.pre[i][j] = fg.pre[k][j]
				}
			}
		}
	}
}
