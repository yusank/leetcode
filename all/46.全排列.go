/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package main

import "log"

func main() {
	a := []int{1, 2, 3}
	r := permute(a)
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
		}
	}
	return res
}

// @lc code=end
