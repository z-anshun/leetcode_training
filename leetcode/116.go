package leetcode

/*
 @Author: as
 @Date: Creat in 1:32 2020/12/29
 @Description: algorithm
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
116.填充每个节点的下一个右侧节点指针
题目：给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点，
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。
如果找不到下一个右侧节点，则将 next 指针设置为 NULL

*/

// 思路:首先，根节点的next一定为null,然后再继续遍历，
// 1.若该节点有右节点，则左节点的next则可以指向右节点
// 2.若该节点没有next（即该节点的右边无同级节点），则该节点的右节点无next；若有，则该节点的右节点的next指向其next的左节点
// 3.递归遍历左子树和右子树
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Next = nil
	modify(root)
	return root
}
func modify(r *Node) {
	if r.Left == nil {
		return
	}
	r.Left.Next = r.Right
	if r.Next != nil {
		r.Right.Next = r.Next.Left
	}
	modify(r.Left)
	modify(r.Right)
}
