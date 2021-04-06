/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	var (
		ys =1
	)

	for i := len(digits)-1;i >= 0;i-- {
		digits[i] += ys
		if digits[i] < 10 {
			return digits
		}

		digits[i]=0
	}

	return append([]int{1}, digits...)


}
// @lc code=end

