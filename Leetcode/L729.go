/**
 * @Author: Gong Yanhui
 * @Description: 729. 我的日程安排表 I
 * @Date: 2025-01-02 22:53
 */

package main

type pair struct {
	start, end int
}

type MyCalendar struct {
	book []pair
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for _, p := range this.book {
		if startTime < p.end && endTime > p.start {
			// 有重叠
			return false
		}
	}
	this.book = append(this.book, pair{startTime, endTime})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

/** 如何判断两个区间是否有重叠
对于两个区间 [a, b) 和 [c, d) 如果没有重叠，那么有 a >= d 或者 b <= c
用反证法证明，如果 a < d 且 b > c，那么这两个区间有重叠
举例3种情况:
左重叠	[a c b d]
右重叠	[c a d b]
包含		[c a b d]
*/
