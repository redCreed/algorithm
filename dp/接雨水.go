package dp

/*
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
	第1种使用动态规划来写

	对于某个柱子arr[ i ]而言，计算它能接的水cap[ i ]其实只需要找它左边的最大值iLeft和右边的最大值iRight，
	这两个最大值就像木桶的两块木板把arr[ i ]夹在中间决定了cap[ i ]的上限，而具体能接到多少则取决于两者的最小值（短板效应），
	即　cap[ i ] = min ( iLeft , iRight  ) - height[i]	每个位置能接的雨水量就是这个位置左右最大值的较小者减去当前的高度。

	因此解题时只需要定义两个数组leftMax和rightMax，从左往右地遍历arr数组，
	找到每个arr[ i ]左边的最大值存入leftMax；再从右往左地遍历arr数组，找到每个arr[ i ]右边的最大值存入rightMax。
	最后再次遍历arr数组计算每个arr[ i ]的接水量，累加起来就是最终结果ret。
*/

func Solution(arr []int) int {
	leftMax := make([]int, len(arr))
	rightMax := make([]int, len(arr))

	leftMax[0] = arr[0]
	//初始化从左往右看的最大值数组
	for i := 1; i < len(arr); i++ {
		leftMax[i] = Max(leftMax[i-1], arr[i])
	}

	rightMax[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], arr[i])
	}

	ret := 0
	for i := 0; i < len(arr); i++ {
		ret += Min(leftMax[i], rightMax[i]) - arr[i]
	}

	return ret
}

// Min 返回较小值
func Min(i, j int) int {
	if i <= j {
		return i
	}

	return j
}
