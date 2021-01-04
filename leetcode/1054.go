package leetcode

import "sort"

/*
 @Author: as
 @Date: Creat in 23:31 2021/1/4
 @Description: algorithm
*/
// 1054. 距离相等的条形码
// 在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。
//
// 请你重新排列这些条形码，使其中两个相邻的条形码 不能 相等。
// 你可以返回任何满足该要求的答案，此题保证存在答案。

//提示：
//1 <= barcodes.length <= 10000
//1 <= barcodes[i] <= 10000

// 思路：排序+填充
type node struct {
	val   int
	count int
}

func rearrangeBarcodes(barcodes []int) []int {
	length := len(barcodes)
	arr := make([]int, length)
	m := make(map[int]int) // 存放数
	for _, v := range barcodes {
		m[v]++
	}
	n := make([]node, 0, length)
	for k, v := range m {
		n = append(n, node{k, v})
	}
	// 根据count排序
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	k := 0 // 初始位置为0
	for _, v := range n {
		val, count := v.val, v.count
		for count > 0 {
			arr[k] = val
			k += 2
			count--
			// 先将偶数位置填好，再填基数位置
			if k >= length {
				k = 1
			}
		}
	}
	return arr
}
