package array

// 双指针原地交换法
func exchange2(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	for left < right { //符合左奇数右偶数的情况
		for left < right && nums[left]%2 == 1 { // left < right 条件避免数组越界
			left++
		}
		for left < right && nums[right]%2 == 0 {
			right--
		} // 不满足左奇数右偶数的情况
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

// 建立空切片，两次遍历
func exchange(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 == 1 {
			res = append(res, num)
		}
	}
	for _, num := range nums {
		if num%2 == 0 {
			res = append(res, num)
		}
	}
	return res
}
