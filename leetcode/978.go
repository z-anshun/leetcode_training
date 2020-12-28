package leetcode

/*
 @Author: as
 @Date: Creat in 1:31 2020/12/29
 @Description: algorithm
*/

// 978. 最长湍流子数组
// 直接遍历查找（动态规划类问题）
func maxTurbulenceSize(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	var count, maxLen int
	var s bool // true 为 > ;false 为 <
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			if !s || count == 0 {
				count++
			} else {
				maxLen = Max(maxLen, count)
				count = 1
			}
			s = true
		} else if arr[i] < arr[i+1] {
			if s || count == 0 {
				count++
			} else {
				maxLen = Max(maxLen, count)
				count = 1
			}
			s = false
		} else {
			maxLen = Max(maxLen, count)
			count = 0
		}
	}
	if count > maxLen {
		return count + 1
	}
	return maxLen + 1
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
