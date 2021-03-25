/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	var (
		m1 = make(map[byte]byte, len(s))
		m2 = make(map[byte]byte, len(s))
	)
	for i := range s {
		tr := t[i]
		if m1[s[i]] > 0 && m1[s[i]] != tr ||
			m2[tr] > 0 && m2[tr] != s[i] {
			return false
		}

		m1[s[i]] = tr
		m2[tr] = s[i]

	}

	return true
}

// @lc code=end

