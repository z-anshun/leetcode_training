package leetcode

import (
	"math"
)

/*
 @Author: as
 @Date: Creat in 16:36 2020/12/28
 @Description: algorithm
*/

/*
题目：在一个小城市里，有 m 个房子排成一排，你需要给每个房子涂上 n种颜色之一（颜色编号为 1 到 n）。
有的房子去年夏天已经涂过颜色了，所以这些房子不需要被重新涂色。

我们将连续相同颜色尽可能多的房子称为一个街区。
比方说 houses = [1,2,2,3,3,2,1,1] ，它包含 5 个街区 [{1}, {2,2}, {3,3}, {2}, {1,1}] 。

给你一个数组houses，一个m * n的矩阵cost和一个整数target，其中：
houses[i]：是第i个房子的颜色，0表示这个房子还没有被涂色。
cost[i][j]：是将第i个房子涂成颜色j+1的花费。

请你返回房子涂色方案的最小总花费，使得每个房子都被涂色后，恰好组成target个街区。
如果没有可用的涂色方案，请返回-1
*/

// 思路：动态规划问题，若前i个房子有k个社区，

// 那么若第i个为j，若第i-1个不为j则前i-1有k-1个社区；
// 若第i-1个为j，那么前i-1个就有k个社区

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	if m < target {
		return -1
	}
	dp := make([][][]int, m) // 定义dp[i][j][k],值为前i个房子，j个社区，最后一个颜色为k的花费
	// 进行初始化
	for i, _ := range dp {
		// 种类最多为m+1，这里为了确保后续正常，所以都为m+1
		dp[i] = make([][]int, m+1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n)
			for k, _ := range dp[i][j] {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	// 如果第一个已经涂过了
	if houses[0] != 0 {
		dp[0][1][houses[0]-1] = 0
	} else {
		// 如果第一个没有涂，那每个都试一试
		for k, _ := range dp[0][1] {
			dp[0][1][k] = cost[0][k]
		}
	}

	// 开始遍历
	for i := 1; i < m; i++ {
		// 最多就 i+1 种
		for j := 1; j <= i+1; j++ {
			// 如果该房子已经涂过了
			if houses[i] != 0 {
				// 相同时
				dp[i][j][houses[i]-1] = dp[i-1][j][houses[i]-1]
				for k := 0; k < n; k++ {
					// 取最小
					// 当最后一个与前一个不同时
					if houses[i]-1 != k {
						dp[i][j][houses[i]-1] = min(dp[i][j][houses[i]-1], dp[i-1][j-1][k])
					}

				}
			} else {
				// 如果没有涂
				for k := 0; k < n; k++ {
					// 最后一个与前一个相同颜色,即前i-1个可以为j个社区
					if dp[i-1][j][k] != math.MinInt32 {
						// 如果前i-1个,j种社区的有值的
						dp[i][j][k] = dp[i-1][j][k] + cost[i][k]
					}
					// 最后一个与前一个不同颜色
					for k1 := 0; k1 < n; k1++ {
						if k1 != k {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][k1]+cost[i][k])
						}
					}

				}

			}
		}
	}

	res := dp[m-1][target][0]
	for _, v := range dp[m-1][target] {
		res = min(res, v)
	}
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
