/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var (
		rs        []int
		left, top int
		right     = len(matrix[0]) - 1
		buttom    = len(matrix) - 1
		counter   = len(matrix) * len(matrix[0])
	)

	for counter > 0 {
		// top
		for i := left; i <= right && counter > 0; i++ {
			rs = append(rs, matrix[top][i])
			counter--
		}
		top++

		for i := top; i <= buttom && counter > 0; i++ {
			rs = append(rs, matrix[i][right])
			counter--
		}
		right--

		for i := right; i >= left && counter > 0; i-- {
			rs = append(rs, matrix[buttom][i])
			counter--
		}
		buttom--

		for i := buttom; i >= top && counter > 0; i-- {
			rs = append(rs, matrix[i][left])
			counter--
		}
		left++
	}

	// [1,2,3,6,9,6,3,2,5]
	return rs
}

// @lc code=end

