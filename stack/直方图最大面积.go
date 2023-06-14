package stack

import "container/list"

func LargestRectangleArea2(heights []int) int {
	//heights := []int{2,1,5,6,2,3}
	// 单调递增栈
	ans := 0
	heights = append([]int{0}, heights...) // 加入左哨兵
	heights = append(heights, []int{0}...) // 加入右哨兵
	stack := []int{0}                      // 栈，左哨兵入栈
	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {
			// 出栈
			popi := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算popi的值
			tmp := heights[popi] * (i - stack[len(stack)-1] - 1)
			if tmp > ans {
				ans = tmp
			}
		}
		// 入栈
		stack = append(stack, i)
	}
	return ans
}

// LargestRectangleArea 给定一个非负数组arr，代表直方图。返回直方图的最大长方形面积。
func LargestRectangleArea(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	//用双向链表代替
	stack := list.New()
	maxArea := 0
	for i := 0; i < len(arr); i++ {
		//单调栈递增  第i个元素小于单调栈的最后一个元素 则单调栈出栈
		for stack.Len() != 0 && arr[i] < arr[stack.Back().Value.(int)] {
			//下表j后面的
			j := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			k := 0
			if stack.Len() == 0 {
				k = -1
			} else {
				k = stack.Back().Value.(int)
			}
			curArea := (i - k - 1) * arr[j]
			maxArea = Max(maxArea, curArea)
		}
		//放在栈顶
		stack.PushBack(i)
	}
	for !(stack.Len() == 0) {
		j := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		k := 0
		if stack.Len() == 0 {
			k = -1
		} else {
			k = stack.Back().Value.(int)
		}
		curArea := (len(arr) - k - 1) * arr[j]
		maxArea = Max(maxArea, curArea)
	}
	return maxArea
}

func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
