/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var (
		minPrice = int(1e9)
		profit   = 0
		max      = func(a, b int) int {
			if a > b {
				return a
			}

			return b
		}
	)

	for _, p := range prices {
		profit = max(profit, p-minPrice)
		if p < minPrice {
			minPrice = p
		}
	}

	return profit
}

// @lc code=end

