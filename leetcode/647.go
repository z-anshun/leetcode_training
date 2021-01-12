package leetcode

/*
 @Author: as
 @Date: Creat in 23:36 2021/1/12
 @Description: algorithm
*/

// 647. 回文子串
// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

func countSubstrings(s string) int {
	length := len(s)
	dp := make([][]bool, length) // 动归
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	res := length
	for right := 0; right < length; right++ {
		for left := 0; left < right; left++ {
			// 左右相等，并且长度小于2或者前面的动归为true
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
				res++
			}
		}
	}
	return res
}
