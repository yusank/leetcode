/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return false
	}

	var (
		sum, maxNum int
	)

	for _, n := range nums {
		sum += n
		if n > maxNum {
			maxNum = n
		}
	}
	
	if sum %2 != 0 {
		return false
	}

	target := sum/2

	if maxNum > target {
		return false
	}

	dp := make([]bool,target+1)
	dp[0]=true

	// [2,10,22,10]
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j]= dp[j] || dp[j-v]
		} 
	}

	return dp[target]
}
// @lc code=end

