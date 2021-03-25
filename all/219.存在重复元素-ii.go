/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return true
		}

		m[nums[i]] = true
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}

	return false
}

// @lc code=end

