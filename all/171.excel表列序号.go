/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	var result int
	for i,r := range columnTitle {
		n := rune(r-'A')
		result += pow(26,len(columnTitle)-1-i)*(int(n)+1)
	}

	return result
}

func pow(base,n int) int {
	var r =1
	for i := 0;i < n;i++ {
		r *= base
	}

	return r
}
// @lc code=end

