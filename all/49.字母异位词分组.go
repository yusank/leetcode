/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	var mp = make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}

		mp[cnt] = append(mp[cnt], str)
	}

	result := make([][]string, 0)
	for _, v := range mp {
		result = append(result, v)
	}

	return result
}

// @lc code=end

