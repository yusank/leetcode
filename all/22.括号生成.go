/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (res []string) {
	var (
		dfs func(string, int, int)
	)

	if n <= 0 {
		return
	}

	dfs = func(str string, l, r int) {
		if l == r && l == 0 {
			res = append(res, str)
			return
		}

		if l == r {
			dfs(str+"(", l-1, r)
			return
		}
		if l < r {
			if l > 0 {
				dfs(str+"(", l-1, r)
			}

			dfs(str+")", l, r-1)
		}
	}

	dfs("", n, n)
	return
}

// @lc code=end

