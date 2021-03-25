/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	temp := make(map[int]bool)
	for n != 1 && !temp[n] {
		temp[n] = true
		n = sum(n)
	}

	return n == 1
}

func sum(n int) int {
	var rs int
	for n != 0 {
		rs += (n % 10) * (n % 10)
		n /= 10
	}

	return rs
}

// @lc code=end

