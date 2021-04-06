/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
 package main

 import "log"

 func main() {
	 merge([]int{0},0,[]int{1},1)
 }

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var (
		i  =  m-1 // nums1 的索引位置
		j = n-1 
		p = m+n-1 // nums1 的最后位置
	)
	for i>= 0 && j >= 0{
		if nums1[i] < nums2[j] {
			nums1[p]=nums2[j]
			j--
		} else {
			nums1[p]= nums1[i]
			i--
		}
		// log.Println(nums1)
		p--
	}

	if i < 0 {
		for j >= 0 {
			nums1[p]=nums2[j]
			p--
			j--
		}
	}

	// log.Println(nums1)
}
// @lc code=end

