/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	var (
		cond, cnt int
	)

	for _,n := range nums {
		if cnt == 0 {
			cond = n
		}

		if cond == n {
			cnt++
		} else {
			cnt--
		}
	}

	return cond
}
// @lc code=end

