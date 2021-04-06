/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	var (
		counter int
		dfs     func([][]byte, int, int)
	)

	dfs = func(g [][]byte, r, c int) {
		nr := len(g)
		nc := len(g[0])

		if r < 0 || c < 0 || r >= nr || c >= nc || g[r][c] == '0' {
			return
		}

		g[r][c] = '0'
		dfs(g, r-1, c)
		dfs(g, r+1, c)
		dfs(g, r, c-1)
		dfs(g, r, c+1)
	}

	nr := len(grid)
	nc := len(grid[0])
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				counter++
				dfs(grid, r, c)
			}
		}
	}

	return counter
}

// @lc code=end

