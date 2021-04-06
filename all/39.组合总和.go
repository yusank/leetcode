/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

func main() {
	a := []int{2, 3, 6, 7}
	t := 7
	r := combinationSum(a, t)
	// log.Println(r)
}

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var (
		combs []int
		ans   [][]int
		dfs   func(int, int)
	)

	dfs = func(target, idx int) {
		if len(candidates) == idx {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, combs...))
			return
		}

		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			combs = append(combs, candidates[idx])
			dfs(target-candidates[idx], idx)
			// log.Println(combs, idx)
			combs = combs[:len(combs)-1]
		}
	}

	dfs(target, 0)
	return ans
}

// @lc code=end
