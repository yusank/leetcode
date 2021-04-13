/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	var (
		low int
		high = len(nums)-1
	)

	for low < high {
		p := low + (high-low)/2

		if nums[p] < nums[high] {
			high = p
		} else {
			low = p+1
		}
	}

	return nums[low]
}
// @lc code=end

