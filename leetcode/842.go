package leetcode

import "math"

/*
 @Author: as
 @Date: Creat in 19:41 2021/1/3
 @Description: algorithm
*/

// 842. 将数组拆分成斐波那契序列
// 给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
//
// 形式上，斐波那契式序列是一个非负整数列表 F，且满足：
//
// 0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
// F.length >= 3；
// 对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
//另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
//
// 返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []

// 思路：属于回溯问题，遍历每一种结果，错误就回溯到上一次的结果
func splitIntoFibonacci(S string) []int {
	length := len(S)
	var F []int
	// 如果第一个为0
	backtrack(S, length, 0, 0, 0, &F)
	return F

}
func backtrack(s string, n, index, sum, prev int, F *[]int) bool {
	// 当遍历到最后一个时
	if index == n {
		return len(*F) >= 3
	}
	cur := 0
	// 循环遍历所有，满足条件退出
	for i := index; i < n; i++ {
		// 除第一个数字外，其它数字不能以0开头
		// 如：“1011” -> [1,0,1,1]
		if i > index && s[index] == '0' {
			break
		}
		// 去 s[index:i+1]的值
		cur = cur*10 + int(s[i]-'0')
		// 确保每个值不大于  2^31 - 1
		if cur > math.MaxInt32 {
			break
		}
		// 如果这个栈已经存了大于两个的数了
		if len(*F) >= 2 {
			// 如果还小，继续取
			if cur < sum {
				continue
			}
			// 如果已经大了，则直接退出
			if cur > sum {
				break
			}
		}
		// 如果这个栈小于两个就直接加，或者满足sun==cur
		*F = append(*F, cur)
		// 循环遍历找后面的，这里的下一个的index为i,sum为当前的数字+上一个的，prev就为当前的数
		if backtrack(s, n, i+1, prev+cur, cur, F) {
			return true
		}
		// 如果不满足条件，回退还原状态，进行下一次循环，更改前一项
		*F = (*F)[:len(*F)-1]
	}
	return false
}
