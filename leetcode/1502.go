package leetcode

/*
 @Author: as
 @Date: Creat in 0:05 2021/1/16
 @Description: algorithm
*/

// 1502. 判断能否形成等差数列
// 给你一个数字数组 arr 。
//
// 如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。
//
// 如果可以重新排列数组形成等差数列，请返回 true ；否则，返回 false 。

func canMakeArithmeticProgression(arr []int) bool {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if min == max {
		return true
	}
	if (max-min)%(len(arr)-1) != 0 {
		return false
	}
	// 算出等差
	sub := (max - min) / (len(arr) - 1)
	tmp := make([]int, len(arr))
	for _, v := range arr {
		// 如果当前数减去最小的不为等差，或者当前数已经存在（若sub!=0，则不能有相同的数存在）
		if (v-min)%sub != 0 || tmp[(v-min)/sub] != 0 {
			return false
		}
		tmp[(v-min)/sub]++
	}
	return true
}
