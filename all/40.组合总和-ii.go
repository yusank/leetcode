/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

package main

import "log"

func main() {
	a := []int{10, 1, 2, 7, 6, 1, 5}
	t := 8

	r := combinationSum2(a, t)
	log.Println(r)
}

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	var (
		combs []int
		ans   [][]int
		dfs   func(int, int)
	)

	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, combs...))
			return
		}

		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			combs = append(combs, candidates[idx])
			dfs(target-candidates[idx], idx+1)
			combs = combs[:len(combs)-1]
		}
	}

	dfs(target, 0)
	return ans
}

// @lc code=end
