package dp

import "fmt"

/*
	动态规划
*/

// Dp 01背包问题 m 背包最大重量
func Dp(m int) {
	//物品重量
	w := []int{1, 4, 3}
	//对应的物品价值
	val := []int{1500, 3000, 2000}
	//表示背包中价值最大值 i j
	f := make([][]int, len(w)+1)
	for i := 0; i <= len(w); i++ {
		f[i] = make([]int, m+1)
	}
	//默认f[0][j]=0 f[i][0]=0
	//i表示放入第i个物品
	for i := 1; i <= len(w); i++ { //不处理第一行
		for j := 1; j <= m; j++ { //不处理第一列  背包中的最大重量
			//当放入第i个商品时
			if w[i-1] > j { //因为时从1开始，所以w[i]改完 w[i-1]
				//f[i][j]表示前i个物品放到最大重量为j的背包的最大价值
				f[i][j] = f[i-1][j]
			} else {
				//公式: f[i][j] = Max(f[i-1][j], val[i]+f[i-1][j-w[i]]) 但是循环都是从1开始所以val[i-1]
				f[i][j] = Max(f[i-1][j], val[i-1]+f[i-1][j-w[i-1]])
			}
		}
	}
	for k, v := range f {
		for k1, _ := range v {
			//打印矩阵
			fmt.Print(f[k][k1], "  ")
		}
		fmt.Print("\n")
	}
}

// Max 比赛a和b的最大值，并返回最大值
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
