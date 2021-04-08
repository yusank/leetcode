/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var (
		ln = len(s)
		start,end , mid int
	)

	if ln < 2 {
		return s
	}

	for i := 0; i<ln;i++ {
		// 单字符为中心夸张
		m1 := maxInRange(s,i,i)
		m2 := maxInRange(s,i,i+1)

		mid = max(mid,max(m1,m2))

		// 找到更大的回文 更新 start end
		// start end 分别赋值 从 当前位置 i 前后加减 mid/2
		if mid > end-start+1 {
			start = i -(mid-1)/2 
			end =mid/2 +i
		}
	}

	return s[start:end+1]
}

// 范围内扩张 查找最长
func maxInRange(str string, s,e int) int {
	l,r := s,e

	for l >=0 && r < len(str) && str[l]==str[r] {
		l--
		r++
	}

	return r-l-1
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
// @lc code=end

