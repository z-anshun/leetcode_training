package leetcode

/*
 @Author: as
 @Date: Creat in 23:16 2021/1/15
 @Description: algorithm
*/

//135. 分发糖果
// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
// 你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
// 那么这样下来，老师至少需要准备多少颗糖果呢？
func candy(ratings []int) int {
	length := len(ratings)
	res := 0
	left := make([]int, length)
	// 从左向右走
	for k, v := range ratings {
		// 如果不是第一个，并且比前一个大
		if k > 0 && v > ratings[k-1] {
			left[k] = left[k-1] + 1
		} else {
			left[k] = 1
		}
	}
	right := 0
	// 从右向左走
	for i := length - 1; i >= 0; i-- {
		// 如果不是最后一个，并且比后一个大
		if i < length-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		// 这里确保了比左右的糖果都多
		res += max(right, left[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
