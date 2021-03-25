/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */
package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	arr := []int{5, 6, 2, 3, 1}
	// sortArray(arr)
	heapSort(arr)
	log.Println(arr)
}

// @lc code=start
func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	randQuickSort(nums, 0, len(nums)-1)
	return nums
}

// 快排
func randQuickSort(nums []int, l, r int) {
	if l < r {
		p := randPartition(nums, l, r)
		randQuickSort(nums, l, p-1)
		randQuickSort(nums, p+1, r)
	}
}

func randPartition(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	p := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= p {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func heapSort(nums []int) {
	buildHeap(nums, len(nums))
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if i == 1 && nums[0] <= nums[i] {
			break
		}

		nums[i], nums[0] = nums[0], nums[i]
		n--
		maxHeapify(nums, 0, i-1)
	}
}

func buildHeap(nums []int, n int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, n)
	}
}

func maxHeapify(nums []int, i, n int) {
	var (
		l = (i << 1) + 1 // *2 +1
		r = (i << 1) + 2
		p int
	)

	if l >= n {
		return
	}

	if r <= n && nums[r] > nums[l] {
		p = r
	} else {
		p = l
	}

	if nums[p] <= nums[i] {
		return
	}

	nums[i], nums[p] = nums[p], nums[i]
	maxHeapify(nums, p, n)
}

// @lc code=end
