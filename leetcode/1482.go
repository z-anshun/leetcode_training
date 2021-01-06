package leetcode

import "math"

/*
 @Author: as
 @Date: Creat in 23:31 2021/1/6
 @Description: algorithm
*/

// 1482. 制作 m 束花所需的最少天数
// 给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
// 现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
// 花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
// 请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1

// m为需要的花束，k为需要的每束花需要的花朵（必须相邻）
func minDays(bloomDay []int, m int, k int) int {
	length := len(bloomDay)

	if length < m*k {
		return -1
	}
	// 最小，最大和返回结果的天数
	min, max, res := math.MaxInt32, 0, 0
	for _, v := range bloomDay {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	l := min
	r := max
	// 这里采用二分法
	for l <= r {
		mid := (r-l)/2 + l
		cur := 0
		m1 := 0
		// 遍历天数
		for i := 0; i < length; i++ {
			// 如果当前的天数小于了mid，则满足
			if bloomDay[i] <= mid {
				cur++
			} else {
				// 反之，相加满足的花束
				m1 += cur / k
				cur = 0
			}
		}
		// 如果全部都满足，则需要再次计算
		m1 += cur / k
		if m1 >= m {
			// 大于或者满足，则赋值res
			r = mid - 1
			res = mid
		} else {
			l = mid + 1
		}
	}
	return res
}
