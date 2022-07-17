/**
 * @Author: Gong Yanhui
 * @Description: 剑指 Offer II 041. 滑动窗口的平均值
 * @Date: 2022-07-17 18:40
 */

package Offer2

type MovingAverage struct {
	size  int
	sum   int
	queue []int
}

// Constructor /** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.size {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	this.sum += val
	this.queue = append(this.queue, val)
	return float64(this.sum) / float64(len(this.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
