package leetcode

/*
 @Author: as
 @Date: Creat in 0:40 2021/1/3
 @Description: algorithm
*/

// 1582. 二进制矩阵中的特殊位置
// 给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，
// 请返回 矩阵 mat 中特殊位置的数目 。
//
// 特殊位置
// 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i, j) 被称为特殊位置。

// 法一：
func numSpecial(mat [][]int) int {
	i := len(mat)    // 列
	n := len(mat[0]) // 行
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	var count int
	for c := 0; c < i; c++ {
		if m2[c] {
			continue
		}
		for r := 0; r < n; r++ {
			if m1[r] {
				continue
			}
			flag := 0
			if mat[c][r] == 1 {
				// 判断这个符不符合要求
				// 列
				for c1 := 0; c1 < i; c1++ {
					if mat[c1][r] != 0 && c1 != c {
						// 这一列不符合要求
						m1[r] = true
						flag = 1
						break
					}
				}
				if flag == 1 {
					break
				}
				// 行
				for r1 := 0; r1 < n; r1++ {
					if mat[c][r1] != 0 && r1 != r {
						// 这一行不符合要求
						m2[c] = true
						flag = 1
						break
					}
				}
				if flag == 0 {
					count++
				}
			}
		}
	}
	return count
}

// 法二：
// 查找特别值,先分别计算行与列的length，再对第i行，或者第j列的合计算
// 最后查找该行该列的合为1的值
func numSpecial2(mat [][]int) int {
	count := 0
	m, n := len(mat), len(mat[0])
	rows, cols := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rows[i] += mat[i][j]
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			cols[j] += mat[i][j]
		}
	}
	for i, v := range rows {
		if v == 1 {
			for j, v2 := range cols {
				// 这里的mat[i][j]很帅
				if v2 == 1 && mat[i][j] == 1 {
					count++
				}
			}
		}
	}
	return count
}
