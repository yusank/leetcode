/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var (
		maxVal int
		temp = make(map[byte]int)
	)

	// dvdf
	for i,j := 0,0; j < len(s);j++ {
		if v,ok := temp[s[j]];ok {
			i = max(v,i)
		}

		maxVal = max(maxVal,j-i+1)
		temp[s[j]] = j+1
	}

	

	return maxVal
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
// @lc code=end

