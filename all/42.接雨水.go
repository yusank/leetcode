/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var (
		p, q                  int // pointer
		partialResult, result int
		ln                    = len(height)
	)

	for p < ln {
		if p == q {
			q++
			continue
		}

	}
}

// @lc code=end

