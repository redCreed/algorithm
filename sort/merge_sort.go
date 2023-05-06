package sort

/*
	归并排序
	https://blog.csdn.net/EDDYCJY/article/details/124185540

	归并排序和快排的相同点：
	1，利用分治思想
	2，具体实现都用递归

	归并排序和快排的不同点：
	1，先分解再合并：归并排序先递归分解到最小粒度，然后从小粒度开始合并排序，自下而上的合并排序；
	2，边分解边排序：快速排序每次分解都实现整体上有序，即参照值左侧的数都小于参照值，右侧的大于参照值；是自上而下的排序；
	3，归并排序不是原地排序，因为两个有序数组的合并一定需要额外的空间协助才能合并；
	4，快速排序是原地排序，原地排序指的是空间复杂度为O(1)；
	5，归并排序每次将数组一分为二，快排每次将数组一分为三

*/

func MergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2

	// 先将左边排序好
	MergeSort(arr, start, mid)

	// 再将右边排序好
	MergeSort(arr, mid+1, end)

	//合并数据
	Merge(arr, start, mid, end)
}

func Merge(arr []int, start, mid, end int) {
	result := make([]int, 0)
	low, high := start, mid+1
	for low <= mid && high <= end {
		if arr[low] > arr[high] {
			//从小到大
			result = append(result, arr[high])
			high++
		} else {
			result = append(result, arr[low])
			low++
		}
	}
	//两子数组长度不一定相同，比较结束后，可能会有一个子数组还有剩余的数没有比较
	if low <= mid {
		result = append(result, arr[low:mid+1]...)
	}

	if high <= end {
		result = append(result, arr[high:end+1]...)
	}
	for k, v := range result {
		arr[start+k] = v
	}
}
