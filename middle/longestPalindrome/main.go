package main

import (
	"fmt"
)

/*
 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	str := "babad"

	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	var (
		l                = len(s)
		start, end, mlen int // mlen 为回文最长长度
	)

	for i := 0; i < l; i++ {
		var (
			// 中心为单个字符串
			l1 = expendaroundcenter(i, i, s)
			// 中心为 2 个字符串 如 abccbada
			l2 = expendaroundcenter(i, i+1, s)
		)

		mlen = max(max(l1, l2), mlen)
		if mlen > end-start+1 {
			// 根据回文长度和当前索引 i 找回文开始和结束位置
			start = i - (mlen-1)/2
			end = i + mlen/2
		}
	}

	return s[start : end+1]
}

//计算以left和right为中心的回文串长度
// babad
func expendaroundcenter(left, right int, s string) int {
	var (
		l = left
		r = right
	)

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
