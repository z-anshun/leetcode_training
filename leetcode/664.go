package leetcode

/*
 @Author: as
 @Date: Creat in 23:13 2021/1/16
 @Description: algorithm
*/

// 664. 奇怪的打印机
// 有台奇怪的打印机有以下两个特殊要求：
//
// 打印机每次只能打印同一个字符序列。
// 每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
// 给定一个只包含小写英文字母的字符串，你的任务是计算这个打印机打印它需要的最少次数。

func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for k, _ := range dp {
		dp[k] = make([]int, n+1)
		dp[k][k] = 1 // 初始化每一段的开头 即[k...k]这一段为1
	}
	// 从长度为2的开始
	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1            // 从[i...i+l-1]这一段
			dp[i][j] = dp[i+1][j] + 1 // 为后一个+1
			for k := i + 1; k <= j; k++ {
				// 如果这段的开头和结尾的相同，则可以单独打印的机会（dp）
				if s[i] == s[k] {
					// [i..j]与[i..k]+[k+1..j]小的那个
					// 这里注意，当k==j时，即[i..j]最后一个与前一个相同时，该段就为[i..j-1]的
					// 确保每一段都是最小的
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
