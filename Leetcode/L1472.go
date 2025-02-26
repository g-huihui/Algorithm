/**
 * @Author: Gong Yanhui
 * @Description: 1472. 设计浏览器历史记录
 * @Date: 2025-02-26 09:50
 */

package main

type BrowserHistory struct {
	homepage string // 默认页面
	history  []string
	index    int // 当前所在页面
}

func BrowserHistoryConstructor(homepage string) BrowserHistory {
	return BrowserHistory{
		homepage: homepage,
		history:  []string{homepage},
		index:    0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.index++
	this.history = this.history[:this.index] // 移除当前页面之后的历史记录
	this.history = append(this.history, url)
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.index {
		steps = this.index
	}
	this.index -= steps

	return this.history[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.index+steps >= len(this.history) {
		this.index = len(this.history) - 1
	} else {
		this.index += steps
	}

	return this.history[this.index]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
