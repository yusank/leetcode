/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

package main

// @lc code=start
func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var (
		direction = "r"
		r, c      int
	)
	for i := 1; i <= n*n; i++ {
		matrix[r][c] = i
		switch direction {
		case "r":
			c++
			if c == n-1-r {
				direction = "d"
			}
		case "l":
			c--
			if c == n-1-r {
				direction = "u"
			}
		case "d":
			r++
			if r == c {
				direction = "l"
			}
		case "u":
			r--
			if r == c+1 {
				direction = "r"
			}
		}
	}

	return matrix
}

// @lc code=end
