package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		minPrefix = strs[0]
		minLen    = len(strs[0])
	)

	for i := 1; i < len(strs); i++ {
		cur := strs[i]
		if len(cur) < minLen {
			minLen = len(cur)
		}

		for j := 0; j < minLen; j++ {
			if cur[j] != minPrefix[j] {
				minPrefix = minPrefix[:j]
				minLen = len(minPrefix)
				break
			}

			if j == minLen-1 {
				minPrefix = minPrefix[:minLen]
			}
		}

		if minLen == 0 {
			minPrefix = ""
			break
		}
	}

	return minPrefix

}
