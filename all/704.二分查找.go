/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// @lc code=end

