/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int,n)
	dp[0]=1
	maxAns := 1

	for i := 1; i < n;i++ {
		dp[i]=1
		for j := 0; j < i;j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i],dp[j]+1)
			}
		}

		maxAns = max(dp[i],maxAns)
	}

	return maxAns
}

func max(a,b int) int {
	if a>b {
		return a
	}

	return b
}
// @lc code=end

