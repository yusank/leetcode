/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
type pair struct{ x, y int }

var directions = []pair{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func exist(board [][]byte, word string) bool {
	var (
		h, w = len(board), len(board[0])
		vis  = make([][]bool, h)
		dfs  func(a, b, c int) bool
	)

	for i := range vis {
		vis[i] = make([]bool, w)
	}

	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		vis[i][j] = true
		defer func() {
			vis[i][j] = false
		}()

		for _, d := range directions {
			if a, b := d.x+i, d.y+j; a >= 0 && a < h && b >= 0 && b < w && !vis[a][b] {
				if dfs(a, b, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, line := range board {
		for j := range line {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

