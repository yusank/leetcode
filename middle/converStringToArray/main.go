package main

import (
	"fmt"
)

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	var (
		s       = "LEETCODEISHIRING"
		numRows = 4
	)

	fmt.Println(convertV2(s, numRows))
}

// 20ms 6.6M
// 第一版本

func convert(s string, row int) string {
	var (
		l = len(s)
	)

	if l <= 1 || row <= 1 {
		return s
	}

	var (
		matrix     = make([][]int, row)
		k          int // string index
		emptyIndex = -1
	)

	for k < l {
		for j := 0; j < row && k < l; j++ {
			arr := matrix[j]
			if emptyIndex == -1 || emptyIndex == j {
				arr = append(arr, k)
				matrix[j] = arr
				k++
			}
		}

		if emptyIndex == -1 && row > 2 {
			emptyIndex = row - 2
			continue
		}

		if emptyIndex > 1 {
			emptyIndex--
			continue
		}

		emptyIndex = -1
	}

	var result string

	for _, arr := range matrix {
		for _, v := range arr {
			str := " "
			if v != -1 {
				str = string(s[v])
				result += str
			}
		}
	}

	return result
}

func convertV2(s string, rowNum int) string {
	if rowNum < 2 {
		return s
	}

	var (
		i, flag = 0, -1
		rows    = make([]string, rowNum)
	)

	for _, r := range s {
		rows[i] += string(r)
		if i == 0 || i == rowNum-1 {
			flag = -flag
		}

		i += flag
	}

	var result string
	for _, str := range rows {
		result += str
	}

	return result
}
