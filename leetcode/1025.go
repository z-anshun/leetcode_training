package leetcode

/*
 @Author: as
 @Date: Creat in 23:47 2021/1/11
 @Description: algorithm
*/

// 1025. 除数博弈
// 爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
//
//最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
//
//选出任一 x，满足 0 < x < N 且 N % x == 0 。
//用 N - x 替换黑板上的数字 N 。
//如果玩家无法执行这些操作，就会输掉游戏。
//
//只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 False。假设两个玩家都以最佳状态参与游戏。

// 数学法
func divisorGame(N int) bool {
	return N%2 == 0
}

// 动规解决问题
func divisorGame2(N int) bool {
	f := make([]bool, N+1)
	f[1], f[2] = false, true
	for i := 3; i < N; i++ {
		for j := 1; j < i; j++ {
			// 如果能整除，并且后手的失败，则为true
			if i%j == 0 && !f[i-j] {
				f[i-j] = true
				break
			}
		}
	}
	return f[N]
}
