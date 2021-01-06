package leetcode

/*
 @Author: as
 @Date: Creat in 23:49 2021/1/5
 @Description: algorithm
*/

// 834. 树中距离之和
// 给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。
// 第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。
// 返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

var (
	sum, count []int
	m          [][]int
)

func sumOfDistancesInTree(N int, edges [][]int) []int {
	count = make([]int, N)
	sum = make([]int, N)
	m = make([][]int, N)
	for i := 0; i < N; i++ {
		count[i] = 1
	}
	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}
	dfsRoot(0, -1)
	dfsAll(0, -1, N)
	return sum
}

func dfsRoot(now, root int) {
	for _, n := range m[now] {
		if n == root {
			continue
		}
		// 这里的root和n是换了的
		// count为深度，sum为该节点所拥有的值
		dfsRoot(n, now)
		count[now] += count[n]        // 与该点连通的点数
		sum[now] += sum[n] + count[n] // 该节点总的值，子节点的总值+连通的点
	}
}

func dfsAll(now, root int, N int) {
	for _, n := range m[now] {
		if n == root {
			continue
		}
		sum[n] = sum[now] - count[n] + N - count[n] // 当前
		dfsAll(n, now, N)
	}
}
