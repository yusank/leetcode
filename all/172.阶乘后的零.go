/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	a,b := solve(n,2),solve(n,5)

	if a > b {
		return b
	}

	return a
}

func solve(n,x int) int {
	var cnt int
	for n > 0 {
		n /= x
		cnt += n
	}

	return cnt
}
// @lc code=end

