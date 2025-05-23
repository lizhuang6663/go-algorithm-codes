package main

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/minimum-path-sum/description/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = minMinPathSum(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func minMinPathSum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
