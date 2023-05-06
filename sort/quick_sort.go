package sort

/*  快速排序
https://blog.csdn.net/weixin_42117918/article/details/81836864?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-0-81836864-blog-124578468.235^v32^pc_relevant_increate_t0_download_v2_base&spm=1001.2101.3001.4242.1&utm_relevant_index=3
*/

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	//选取数组第一个元素作为比较值，将小于该值的数放到左边，大于该值放到右边
	value := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		//小于对比值，放左边
		if arr[i] <= value {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}

	sortedLow := QuickSort(low)
	sortedHigh := QuickSort(high)
	sorted := make([]int, 0)
	sorted = append(sorted, sortedLow...)
	sorted = append(sorted, value)
	sorted = append(sorted, sortedHigh...)
	return sorted
}
