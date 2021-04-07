/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
 package main

 import "log"

 func main() {
	t := "Today is a nice day "
	log.Println(lengthOfLastWord(t))
 }

// @lc code=start
func lengthOfLastWord(s string) int {
	var (
		c int
	)
	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' {
			c++
			continue
		}
		
		// 如果为空
		if c != 0 {
			break
		}
	}

	return c
}
// @lc code=end

