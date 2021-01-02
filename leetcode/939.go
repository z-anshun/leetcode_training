package leetcode

import "math"

/*
 @Author: as
 @Date: Creat in 0:36 2021/1/3
 @Description: algorithm
*/

//939. 最小面积矩形
//给定在 xy 平面上的一组点，确定由这些点组成的矩形的最小面积，其中矩形的边平行于 x 轴和 y 轴。
//
//如果没有任何矩形，就返回 0。

// 最小矩阵
// 因为length < 40000,于是便调用map标记，看这个数是否存在，再遍历，根据对角查看
// O(n^2)
func minAreaRect(points [][]int) int {
	// length := len(points)
	m := make(map[int]bool, len(points))
	for _, v := range points {
		m[v[0]*1e5+v[1]] = true
	}
	size := math.MaxInt32
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for k := 0; k < len(points); k++ {
			x2, y2 := points[k][0], points[k][1]
			if x1 == x2 || y2 == y1 {
				continue
			}
			now := (x1 - x2) * (y2 - y1)
			if now > 0 && m[1e5*x2+y1] && m[1e5*x1+y2] && now < size {
				size = now

			}

		}

	}

	if size == math.MaxInt32 {
		return 0
	}
	return size
}
