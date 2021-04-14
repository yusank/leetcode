/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	var (
		n, m = len(text1), len(text2)
		dp   = make([][]int, n+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[n][m]

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

