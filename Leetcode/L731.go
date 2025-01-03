/**
 * @Author: Gong Yanhui
 * @Description: 731. 我的日程安排表 II
 * @Date: 2025-01-03 12:52
 */

package main

type pairs struct {
	start, end int
}

type MyCalendarTwo struct {
	books []pairs // 保存所有的预定时间
	over  []pairs // 保存所有一次重叠时间
}

func MyCalendarTwoConstructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	// 现在已经重叠的数据里面找是否有重叠
	for _, p := range this.over {
		if startTime < p.end && endTime > p.start {
			return false
		}
	}

	// 现在已经预定的数据里面找是否有重叠
	for _, p := range this.books {
		if startTime < p.end && endTime > p.start {
			// 有重叠，将重叠的时间加入到over中
			this.over = append(this.over, pairs{max(startTime, p.start), min(endTime, p.end)})
		}
	}

	this.books = append(this.books, pairs{startTime, endTime})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

func main() {
	c := MyCalendarTwoConstructor()
	println(c.Book(10, 20)) // true
	println(c.Book(50, 60)) // true
	println(c.Book(10, 40)) // true
	println(c.Book(5, 15))  // false
	println(c.Book(5, 10))  // true
	println(c.Book(25, 55)) // true
}
