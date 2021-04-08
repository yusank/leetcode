/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package main

import "log"

func main() {
	a := []int{1, 2, 3}
	r := permute2(a)
	log.Println(r)
}

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
			log.Println(res)
		}
	}
	return res
}

// 回溯法 2
func permute2(nums []int) [][]int {
	var (
		rs [][]int
		dfs func([]int)
	)

	dfs = func(trace []int) {
		if len(trace) == len(nums) {
			rs = append(rs, append([]int{},trace...))
			return
		}

		for i := range nums {
			var ex bool
			for _,t := range trace {
				if nums[i] == t {
					ex=true
					break
				}
			}

			if ex {
				continue
			}

			trace = append(trace,nums[i])
			dfs(trace)
			trace = trace[:len(trace)-1]
		}
	}

	dfs([]int{})

	return rs
}

// @lc code=end
