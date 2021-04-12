/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	var (
		ans []string
		cnt int
		str string
	)

	for i,b := range s {
		if b == ' ' {
			if cnt == 0 {
				continue
			}

			ans = append(ans, s[i-cnt:i])
			cnt = 0
			continue
		}

		cnt++
	}

	if cnt != 0 {
		ans = append(ans,s[len(s)-cnt:])
	}

	for i := len(ans)-1;i >= 0;i-- {
		str += ans[i]
		if i > 0 {
			str += " "
		}
	}

	return str
}
// @lc code=end

