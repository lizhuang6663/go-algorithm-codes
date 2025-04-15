package main

// 动态规划 一、入门DP 1.1 爬楼梯
// 力扣：https://leetcode.cn/problems/combination-sum-iv/description/

// 本质是爬楼梯，相当于每次往上爬 nums[i] 步
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}

		// nums = [1,2,3], target = 4
		// 输出：7
		// dp[1] = dp[0]
		// dp[2] = dp[1] + dp[0]
		// dp[3] = dp[2] + dp[1] + dp[0]

	}
	return dp[target]
}
