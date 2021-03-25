package main

import "fmt"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (36.31%)
 * Likes:    894
 * Dislikes: 0
 * Total Accepted:    128.3K
 * Total Submissions: 353.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须 原地 修改，只允许使用额外常数空间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[1,2,3]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,5]
 * 输出：[1,5,1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */


func main() {
	t := []int{4,5,2,6,3,1}
	nextPermutation(t)

	fmt.Println(t)
}

func nextPermutation(nums []int)  {
	var (
		n = len(nums)
		i = n-2
	)

	// 找到较小的数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n-1
		// 找到较大的数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}

		nums[i],nums[j]=nums[j],nums[i]
	}	

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2;i++ {
		nums[i],nums[len(nums)-1-i] = nums[len(nums)-1-i],nums[i]
	}
}