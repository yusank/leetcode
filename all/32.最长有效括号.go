/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	var (
		ml          = 0
		left, right int
	)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			ml = max(ml, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}

		if left == right {
			ml = max(ml, 2*left)
		} else if right < left {
			left, right = 0, 0
		}
	}

	return ml
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

