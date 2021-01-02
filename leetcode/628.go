package leetcode

import (
	"math"
	"sort"
)

/*
 @Author: as
 @Date: Creat in 23:36 2021/1/2
 @Description: algorithm
*/

// 628. 三个数的最大乘积
// 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积

// 注意: 给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
// 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围

// 法一：排序
func maximumProduct1(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	if nums[0] >= 0 || nums[1] >= 0 {
		return nums[length-3] * nums[length-2] * nums[length-1]
	}
	max1 := nums[0] * nums[1] * nums[len(nums)-1]
	max2 := nums[length-3] * nums[length-2] * nums[length-1]
	if max2 > max1 {
		return max2
	}
	return max1
}

// 法二：线性查找
// 很明显，只用找到3个最大的，2个最小的就行了
func maximumProduct(nums []int) int {
	var min1, min2 int
	var max1, max2, max3 int
	min1 = math.MaxInt32
	min2 = min1
	max1 = math.MinInt32
	max3 = max1
	max2 = max1
	for _, v := range nums {
		if v < min1 {
			min1 = v

		} else if v < min2 {
			min2 = v
			min1 = min2

		}
		if max1 < v {
			max3 = max2
			max2 = max1
			max1 = v
		} else if max2 < v {
			max3 = max2
			max2 = v

		} else if max3 < v {
			max1 = v
		}
	}

	if min1*min2*max1 >= max1*max2*max3 {
		return min1 * min2 * max1
	} else {
		return max1 * max2 * max3
	}
}
