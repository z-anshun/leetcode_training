package leetcode

/*
 @Author: as
 @Date: Creat in 23:31 2021/1/18
 @Description: algorithm
*/

//467. 环绕字符串中唯一的子字符串
// 把字符串 s 看作是“abcdefghijklmnopqrstuvwxyz”的无限环绕字符串，所以 s 看起来是这样的："...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
//
// 现在我们有了另一个字符串 p 。你需要的是找出 s 中有多少个唯一的 p 的非空子串，尤其是当你的输入是字符串 p ，你需要输出字符串 s 中 p 的不同的非空子串的数目。
//
// 注意: p 仅由小写的英文字母组成，p 的大小可能超过 10000。

// 相当于寻找连续的子字符串
func findSubstringInWraproundString(p string) int {
	length := len(p)
	if length <= 1 {
		return length
	}
	var dp [26]int
	cnt := 0
	for i := 0; i < length; i++ {
		if i > 0 &&
			(p[i] == p[i-1]+1 || p[i-1] == p[i]+25) {
			cnt++
		} else {
			cnt = 1
		}
		// 以当前为结尾的最长子字符串
		dp[p[i]-'a'] = max(cnt, dp[p[i]-'a'])
	}
	res := 0
	for _, v := range dp {
		// 所有以当前结尾的子字符串总和 就为长度
		// 如"abc" 中，以'c'结尾的子字符串有："abc","bc","c"
		res += v
	}
	return res
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
