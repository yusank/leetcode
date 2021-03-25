/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
package main

import (
	"log"
)

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	log.Println(isValidSudoku(input))
}

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for i, line := range board {
		if !isValidSlice(line) {
			return false
		}

		// 第一行检查每一列
		if i == 0 {
			for jj := 0; jj < 9; jj++ {
				slice := make([]byte, 0)
				for ii := 0; ii < 9; ii++ {
					slice = append(slice, board[ii][jj])
				}

				if !isValidSlice(slice) {
					return false
				}
			}
		}

		// 每三行检查这一行的三个3x3小格子
		if i%3 == 0 {
			slice := make([]byte, 0)
			for jj := 0; jj < 9; jj++ {
				for ii := 0; ii < 3; ii++ {
					slice = append(slice, board[i+ii][jj])
				}

				if (jj+1)%3 == 0 {
					if !isValidSlice(slice) {
						return false
					}
					slice = make([]byte, 0)
				}
			}
		}
	}

	return true
}

func isValidSlice(slice []byte) bool {
	temp := make(map[byte]bool)
	for _, b := range slice {
		if b == byte('.') {
			continue
		}

		if temp[b] {
			return false
		}

		temp[b] = true
	}

	return true
}

// @lc code=end
