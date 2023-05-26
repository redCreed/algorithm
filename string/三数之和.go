package string

import (
	"fmt"
	"sort"
)

/*
	给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与
	target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

	nums := []int{-1,2,1,4} target=4
	输出:2
	与target最接近的和是2(-1+2+1=2)
*/

// ViolenceThreeSum 采用暴力遍历 3层for循环处理
func ViolenceThreeSum(nums []int, target int) {
	//求出3个数之和
	sums := make([]int, 0)
	//此处i的范围
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				//遍历求3个数之和
				sums = append(sums, nums[i]+nums[j]+nums[k])
			}
		}
	}
	fmt.Println("sums:", sums)
	//定义最接近的值
	min := 10000
	data := 0
	for i := 0; i < len(sums); i++ {
		if abs(sums[i]-target) < min {
			min = abs(sums[i] - target)
			data = sums[i]
		}
	}

	fmt.Println("最接近的值是:", data)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
	排序+双指针
	我们使用双指针的思想，我们把三层循环，改为一层循环嵌套一个while的双指针来判断
	我们让start=i+1 end=nums.size()-1
	并且设置ans的初值为数组中的前三个数
	我们在遍历的同时比较现在三数之和距离target的距离和ans和target的距离
*/

func PointerThreeSum(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	//排序从小到大
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			current := nums[i] + nums[left] + nums[right]
			if abs(sum-target) > abs(current-target) {
				sum = current
			}

			if sum < target {
				left++
			} else if sum == target {
				return sum
			} else {
				right--
			}
		}

	}
	return sum
}
