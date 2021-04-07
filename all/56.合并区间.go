/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	var rs [][]int

	// 按每组的 start 升序排序
	sort.Slice(intervals,func(i,j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0;i < len(intervals);i++{
		prev := intervals[i]

		// 判断 结果集最后组的 end 与当前遍历到的组的 end
		if len(rs) == 0 || rs[len(rs)-1][1] < prev[0] {
			rs = append(rs,[]int{prev[0],prev[1]})
		} else {
			rs[len(rs)-1][1] = max(rs[len(rs)-1][1], prev[1])
		}
		
	}

	return rs
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
// @lc code=end

