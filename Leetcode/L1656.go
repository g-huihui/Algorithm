/**
 * @Author: Gong Yanhui
 * @Description: 1656. 设计有序流
 * @Date: 2025-02-24 23:29
 */

package main

type OrderedStream struct {
	index int      // 指向当前应该返回的位置
	size  int      // 数据长度
	data  []string // 数据
}

func OrderedStreamConstructor(n int) OrderedStream {
	return OrderedStream{
		index: 0,
		size:  n,
		data:  make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	var res = make([]string, 0)
	if idKey > this.size || idKey < 0 { // 超出范围
		return res
	}
	// 插入数据
	this.data[idKey-1] = value
	// 查找连续的数据
	for this.index < this.size && this.data[this.index] != "" {
		res = append(res, this.data[this.index])
		this.index++
	}

	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
