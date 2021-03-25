/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	var (
		result string
		n = 26
	)
	for columnNumber != 0 {
		columnNumber--
		result = string(columnNumber % n + 'A')+result
		columnNumber /= n
	}

	return result
}


// @lc code=end

