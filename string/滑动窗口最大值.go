package string

/*
	给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
	在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。
*/

// ViolenceMaxSlidingWindow 暴力解法
func ViolenceMaxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		//倒数第k个时
		if i > len(nums)-k {
			return ret
		}

		tempMax := 0
		for j := i; j < i+k; j++ {
			if tempMax < nums[j] {
	 			tempMax = nums[j]
			}
		}
		ret = append(ret, tempMax)
	}

	return ret
}

// QueueMaxSlidingWindow 双端队列：是一种具有队列和栈的性质的数据结构。双端队列中的元素可以从两端弹出或者插入。
func QueueMaxSlidingWindow() {

}
