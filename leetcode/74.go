package leetcode

/*
 @Author: as
 @Date: Creat in 1:27 2020/12/30
 @Description: algorithm
*/

// 74. 搜索二维矩阵
// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。

// 有序，就二分
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	}

	low := 0
	high := m*n - 1
	mid := (high-low)/2 + low

	for low < high && mid != 0 {
		r := mid / n // 行
		c := mid % n // 列
		if matrix[r][c] > target {
			high = mid - 1
		} else if matrix[r][c] < target {
			low = mid + 1
		} else {
			return true
		}
		mid = (high-low)/2 + low
	}
	// 因为low为mid+1，当mid=high时，就会超过index
	// 前面确保low有效。后面是因为最后一次，可能会没比较，所以最后要验证
	if low >= m*n || matrix[low/n][low%n] != target {
		return false
	}
	return true
}
