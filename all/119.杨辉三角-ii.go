/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	row1 := []int{1, 1}
	for i := 1; i < rowIndex; i++ {
		row1 = nextRow(row1)
	}

	return row1
}

func nextRow(last []int) []int {
	next := make([]int, len(last)+1)
	next[0] = 1
	next[len(next)-1] = 1

	for i := 1; i < len(next)-1; i++ {
		next[i] = last[i] + last[i-1]
	}

	return next
}

// @lc code=end

