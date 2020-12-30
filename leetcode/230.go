package leetcode

/*
 @Author: as
 @Date: Creat in 23:45 2020/12/30
 @Description: algorithm
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. 二叉搜索树中第K小的元素
// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

// 思路:二叉树中，序已经排好了，左节点<根节点<右节点
// 因此，先遍历左节点即可
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		// 如果该节点不为nil,则直接加入，再遍历左节点后，得到一个值从大到小的栈
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			// 如果该节点为空了，即遍历左节点到尾了
			// 并且有节点
		} else if l := len(stack); l != 0 {
			last := stack[l-1]
			if k == 1 {
				return last.Val
			}
			k--
			// 后面的节点相当于递归，去除最后一个最小的（因为k--,这里就要去除最小的）
			root = last.Right // 寻找右节点，若没有，则继续判断最后一个节点；若有，则添加并寻找其左节点
			stack = stack[:l-1]
		}
	}
}
