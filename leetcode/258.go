package leetcode

/*
 @Author: as
 @Date: Creat in 21:24 2021/1/7
 @Description: algorithm
*/

// 258. 各位相加
// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

// 思路：满9 实现+1，不满就保持；直到为 <10.换句话说，就是取9的模
func addDigits(num int) int {
	// 这里num-1是为了避免num%9==0的情况
	return (num-1)%9 + 1
}
