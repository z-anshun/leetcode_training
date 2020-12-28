package leetcode

import "math"

/*
 @Author: as
 @Date: Creat in 8:12 2020/12/14
 @Description: algorithm
*/

// 581. 最短无序连续子数组
//	给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，
//	那么整个数组都会变为升序排序。
//	你找到的子数组应是最短的，请输出它的长度。
//  思路：线性动规类问题，
//  从左往右遍历找最大值，从右往左遍历找最小值（如果不满足升序或者降序，就对high或low赋值）
func findUnsortedSubarray(nums []int) int {
	var high, low int
	min := math.MaxInt32
	max := math.MinInt32
	// 从左向右遍历,找到不满足升序的最高点
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			high = i
		}
	}
	// 从右向左，找到不满足降序的最低点
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			low = i
		}
	}
	if high <= low {
		return 0
	}
	return high - low + 1
}
