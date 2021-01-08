package leetcode

/*
 @Author: as
 @Date: Creat in 23:17 2021/1/8
 @Description: algorithm
*/

// 1247. 交换字符使得字符串相同
// 有两个长度相同的字符串 s1 和 s2，且它们其中 只含有 字符 "x" 和 "y"，
// 你需要通过「交换字符」的方式使这两个字符串相同。
// 每次「交换字符」的时候，你都可以在两个字符串中各选一个字符进行交换。
// 交换只能发生在两个不同的字符串之间，绝对不能发生在同一个字符串内部。
// 也就是说，我们可以交换 s1[i] 和 s2[j]，但不能交换 s1[i] 和 s1[j]。
// 最后，请你返回使 s1 和 s2 相同的最小交换次数，如果没有方法能够使得这两个字符串相同，则返回 -1 。

// 提示：
// 1 <= s1.length, s2.length <= 1000
// s1, s2 只包含 'x' 或 'y'。
func minimumSwap(s1 string, s2 string) int {
	// 记录不同的时候即可
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return -1
	}
	var cntX, cntY int // 相同位置下，s1为x，s2为y的次数和s1为y，s2为x的次数
	for i := 0; i < l1; i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			cntX++
		} else if s1[i] == 'y' && s2[i] == 'x' {
			cntY++
		}
	}
	// 肯定不能为奇数
	if (cntX+cntY)%2 == 1 {
		return -1
	}
	// 偶数，每两个cntX或cntY换一次，每个cntX和cntY换两次
	return cntX/2 + cntY/2 + cntX%2*2
}
