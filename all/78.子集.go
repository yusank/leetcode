/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	fmt.Println(subsets(input))
}

// @lc code=start
func subsets(nums []int) [][]int {
	var (
		ans  [][]int
		dfs  func(int)
		temp []int
	)

	dfs = func(idx int) {
		if len(nums) == idx {
			ans = append(ans, append([]int{}, temp...))
			return
		}

		temp = append(temp, nums[idx])
		fmt.Println(idx, temp)
		dfs(idx + 1)
		temp = temp[:len(temp)-1]
		dfs(idx + 1)

	}

	dfs(0)
	return ans
}

// @lc code=end
