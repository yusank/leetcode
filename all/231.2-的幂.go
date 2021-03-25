/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */
package main

func main() {
	isPowerOfTwo(8)
}

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// @lc code=end
