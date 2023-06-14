package graph

func MaxIslandArea(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ans = Max(ans, dfs(i, j, grid))
		}
	}

	return ans
}

func dfs(i, j int, grid [][]int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	//当前点为1 让i j等于0表示已经访问 同时默认面积为1
	grid[i][j] = 0
	ans := 1
	//向左一位
	ans += dfs(i-1, j, grid)
	//向右一位
	ans += dfs(i+1, j, grid)
	//向上一位
	ans += dfs(i, j-1, grid)
	//向下一位
	ans += dfs(i, j+1, grid)

	return ans
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
