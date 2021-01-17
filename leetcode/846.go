package leetcode

import (
	"sort"
)

/*
 @Author: as
 @Date: Creat in 23:52 2021/1/17
 @Description: algorithm
*/
// 846. 一手顺子
// 爱丽丝有一手（hand）由整数数组给定的牌。
//
// 现在她想把牌重新排列成组，使得每个组的大小都是 W，且由 W 张连续的牌组成。
//
// 如果她可以完成分组就返回 true，否则返回 false。

func isNStraightHand(hand []int, W int) bool {
	length := len(hand)
	if length%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}
	sort.Ints(hand)
	cnt, cur := 0, 0
	m := make(map[int]bool, length)
	for i := 0; i < length; i++ {
		if m[i] {
			continue
		}
		m[i] = true
		cur, cnt = i, 1
		// 从i的后面开始遍历
		for j := i + 1; j < length; j++ {
			// 如果被标记了，就继续
			if m[j] {
				continue
			}
			// 如果当前的 - 上一个的大于1，也就是不连续了，就返回false
			if hand[j]-hand[cur] > 1 {
				return false
			} else if hand[j]-hand[cur] == 1 {
				// 如果连续
				m[j] = true // 标记已被使用
				cnt++       // 数组+1
				cur = j     // 记录当前位置
				if cnt == W {
					break
				}
			}
		}
	}
	if cnt != W {
		return false
	}
	return true

}
