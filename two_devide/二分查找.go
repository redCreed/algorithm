package two_devide

/*
	二分查找的数组必须有序
*/

// TwoDevide 查找指定数的下表
func TwoDevide(arr []int, target int) int {
	if len(arr) <= 1 {
		if arr[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := len(arr) - 1
	mid := 0
	for left <= right {
		mid = (right + left) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}

	return -1
}
