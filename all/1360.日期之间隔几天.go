/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */

// @lc code=start
func daysBetweenDates(date1 string, date2 string) int {
	year1, _ := strconv.Atoi(date1Arr[0：3])
	month1, _ := strconv.Atoi(date1Arr[3：5])
	day1, _ := strconv.Atoi(date1Arr[5：])
	year2, _ := strconv.Atoi(date2Arr[0：3])
	month2, _ := strconv.Atoi(date2Arr[3：5])
	day2, _ := strconv.Atoi(date2Arr[5：])

	sub：= dayDistance(year1, month1, day1) - dayDistance(year2, month2, day2)
	if sub >= 0 {
		return sub
	}

	return -sub
}

func dayDistance(year, month, day int) int {
	monthDays := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// 当月的天数差
	res := day - 1
	// 每月的天数差
	for month != 0 {
		month--
		res += monthDays[month]
		if month == 2 && isLeapYear(year) {
			res += 1
		}
	}

    // 每年的天数差
	res += 365 * (year - 1971)
	// 下面是为了解决闰年多一天的问题
    res += (year-1)/4 - 1971/4
    res -= (year-1)/100 - 1971/100
    res += (year-1)/400 - 1971/400
	
	return res
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
// @lc code=end

