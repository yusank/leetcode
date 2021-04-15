/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
package main

func main() {
	grid := [][]int{
		{1,2,3},
		{4,5,6},
	}

	minPathSum(grid)
}

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	var (
		dp = make([][]int,len(grid))
	)

	for i := 0;i < len(dp);i++ {
		dp[i]=make([]int,len(grid[0]))
	}

	dp[0][0] = grid[0][0]

	for i :=1; i < len(grid[0]);i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i :=1; i < len(grid);i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid);i++ {
		for j := 1; j  <  len(grid[0]);j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j],dp[i][j-1])
		}
	}
	
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a,b int) int {
	if a>b {
		return b
	}

	return a
}
// @lc code=end

