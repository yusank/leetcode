/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	result := make([][]int, numRows)
	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		result[i] = nextRow(result[i-1])
	}

	return result
}

func nextRow(last []int) []int {
	next := make([]int, len(last)+1)
	next[0] = 1
	next[len(next)-1] = 1
	for i := 1; i < len(next)-1; i++ {
		next[i] = last[i-1] + last[i]
	}

	return next
}

// @lc code=end

