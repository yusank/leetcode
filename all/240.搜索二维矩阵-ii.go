/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
		{11, 13, 15, 17, 19},
		{12, 14, 16, 18, 20},
		{21, 22, 23, 24, 25},
	}

	fmt.Println(searchMatrix2(matrix, 21))
}

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	col := 0
	row := len(matrix) - 1

	for col < len(matrix[0]) && row >= 0 {
		// fmt.Println("1")
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			row--
		} else {
			col++
		}

	}

	return false
}

// @lc code=end
