/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	var s, e int = -1, -1
	for i, n := range nums {
		if n != target {
			continue
		}

		if s == -1 {
			s = i
		}
		e = i
	}

	return []int{s, e}

}

// @lc code=end

