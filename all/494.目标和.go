/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	var (
		dp = make([][]int, len(nums))
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2001)
	}

	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] += 1

	for i := 1; i < len(nums); i++ {
		for j := -1000; j <= 1000; j++ {
			if dp[i-1][j+1000] > 0 {
				dp[i][j+nums[i]+1000] += dp[i-1][j+1000]
				dp[i][j-nums[i]+1000] += dp[i-1][j+1000]
			}
		}
	}

	if target > 1000 {
		return 0
	}

	return dp[len(nums)-1][target+1000]

}

// @lc code=end

