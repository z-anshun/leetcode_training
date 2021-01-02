package leetcode

/*
 @Author: as
 @Date: Creat in 0:21 2021/1/3
 @Description: algorithm
*/
// 198. 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。
// 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你
// 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 典型的动态规划问题
func rob(nums []int) int {

	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	max := make([]int, length) // 记录其前i+1个的最大值
	max[0] = nums[0]
	if nums[0] > nums[1] {
		max[1] = nums[0]
	} else {
		max[1] = nums[1]
	}
	if length == 2 {
		return max[1]
	}
	max[2] = nums[2] + nums[0]
	for k := 3; k < length; k++ {
		if max[k-2] < max[k-3] {
			max[k] = max[k-3] + nums[k]
		} else {
			max[k] = max[k-2] + nums[k]
		}
	}
	if max[length-1] > max[length-2] {
		return max[length-1]
	}
	return max[length-2]
}
