package stack

import "container/list"

//给定一个非负数组arr，代表直方图。返回直方图的最大长方形面积。
func LargestRectangleArea(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	//用双向链表代替
	stack := list.New()
	maxArea := 0
	for i:=0; i< len(arr); i++ {
		//假如是递增的 stack最后一个元素
		for stack.Len() != 0 && arr[i] < arr[stack.Back().Value.(int)] {
			j := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			k := 0
			if stack.Len() == 0 {
				k=-1
			}else {
				k = stack.Back().Value.(int)
			}
			curArea := (i-k-1)* arr[j]
			maxArea= Max(maxArea, curArea)
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

func Max(a, b int) int  {
	if a >= b {
		return a
	}

	return b
}
