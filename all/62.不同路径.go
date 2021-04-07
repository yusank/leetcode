/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int,m)
	for i := range dp {
		dp[i]=make([]int,n)
	}

	dp[0][0]=1

	for i := 0;i < m;i++ {
		for j := 0; j < n;j++ {
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 {
				dp[i][j]=dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
// @lc code=end

