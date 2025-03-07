package sort

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		// 保证 0 ~ i-1 有序
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		// 填坑 插入位置
		nums[j+1] = key
	}
	return nums
}
