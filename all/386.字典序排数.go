/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	var (
		ans = make([]int, 0)
		dfs func(int)
	)

	dfs = func(i int) {
		if i > n {
			return
		}

		for j := 0; j < 10; j++ {
			if i+j <= n && i+j != 0 {
				ans = append(ans, i+j)
				dfs(10 * (i + j))
			}
		}
	}

	dfs(0)
	return ans
}

// @lc code=end

