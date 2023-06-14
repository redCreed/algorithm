package array

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	i, j := 0, 0
	ret := make([]int, 0)
	for i < len1 && j < len2 {
		if nums1[i] > nums2[j] {
			ret = append(ret, nums2[j])
			j++
		} else {
			ret = append(ret, nums1[i])
			i++
		}
	}
	if i < len1 {
		ret = append(ret, nums1[i:]...)
	}

	if j < len2 {
		ret = append(ret, nums2[j:]...)
	}
	fmt.Println(ret)

	//判断是否偶数
	ll := len(ret)
	if ll%2 == 0 {
		index := ll / 2
		index--
		return float64(ret[index]+ret[index+1]) / 2
	} else {
		index := (ll + 1) / 2
		index--
		return float64(ret[index])
	}

}
