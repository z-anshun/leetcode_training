package leetcode

/*
 @Author: as
 @Date: Creat in 23:42 2021/1/9
 @Description: algorithm
*/

// 796. 旋转字符串
// 给定两个字符串, A 和 B。
//
// A 的旋转操作就是将 A 最左边的字符移动到最右边。
// 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。
// 如果在若干次旋转操作之后，A 能变成B，那么返回True。

func rotateString(A string, B string) bool {
	// 都为空就返回true
	if len(A) == 0 && len(B) == 0 {
		return true
	} else if len(A) != len(B) {
		return false
	}
	for i := 1; i < len(A); i++ {
		// 当A的后缀等于B的前缀的时候
		if A[i:] == B[:len(A)-i] {
			// 匹配A的前缀是否等于B的后缀
			if A[:i] == B[len(A)-i:] {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
