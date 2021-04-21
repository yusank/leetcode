/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0] // dp0 为第i天不持有股票时的收益 dp1 为持有股票时的收益
	for i := 0; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}

	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

