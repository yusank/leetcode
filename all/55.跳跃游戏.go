/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	var m int

	for i := 0; i < len(nums); i++ {
		if i <= m {
			m = max(m, nums[i]+i)
			if m >= len(nums)-1 {
				return true
			}
		}
	}

	return false

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

