/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	var (
		count, pre int
		m          = map[int]int{
			0: 1,
		}
	)

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		count += m[pre-k]
		m[pre] += 1
	}

	return count
}

// @lc code=end

