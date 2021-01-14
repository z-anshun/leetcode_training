package leetcode

/*
 @Author: as
 @Date: Creat in 0:07 2021/1/15
 @Description: algorithm
*/

// 面试题 02.03. 删除中间节点
// 实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// 改变值
	//node.Val=node.Next.Val
	//node.Next=node.Next.Next
	// 改变指针
	*node = *node.Next
}
