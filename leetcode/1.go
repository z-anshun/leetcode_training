package leetcode

/*
 @Author: as
 @Date: Creat in 0:37 2021/1/3
 @Description: 梦开始的地方
*/

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums)-1; i++ {
		index, ok := m[target-nums[i]] // 寻找对应的值是否拥有
		if ok && index != i {
			return []int{i, index}
		}
	}
	return nil
}
