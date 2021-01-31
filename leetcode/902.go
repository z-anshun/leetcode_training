package leetcode

import (
	"math"
	"strconv"
)

/*
 @Author: as
 @Date: Creat in 20:44 2021/1/31
 @Description: algorithm
*/

// 902. 最大为 N 的数字组合
// 我们有一组排序的数字 D，它是  {'1','2','3','4','5','6','7','8','9'} 的非空子集。
//（请注意，'0' 不包括在内。）
//
// 现在，我们用这些数字进行组合写数字，想用多少次就用多少次。
// 例如 D = {'1','3','5'}，我们可以写出像 '13', '551', '1351315' 这样的数字。

// 返回可以用 D 中的数字写出的小于或等于 N 的正整数的数目。
func atMostNGivenDigitSet(digits []string, n int) int {
	str := strconv.Itoa(n)

	length := len(str)
	dp := make(map[int]int, length+1)
	dp[0] = 1 // 这里保证相等的时候可以+1
	for i := length - 1; i >= 0; i-- {
		k := length - i - 1 // 表示第几位，这里我们从第一位开始
		for _, d := range digits {
			// 如果第一个比对应的小，即后面的几位可任意选取
			if d[0] < str[i] {
				dp[k+1] += int(math.Pow(float64(len(digits)), float64(k)))
			} else if d[0] == str[i] {
				// 反之就找它的对应的下一个k
				dp[k+1] += dp[k]
			}
		}
	}
	// 加上长度小于length的数目
	for i := 1; i < length; i++ {
		dp[length] += int(math.Pow(float64(len(digits)), float64(i)))
	}
	return dp[length]
}
