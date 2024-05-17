/**
 * @Author: Gong Yanhui
 * @Description: 380. O(1) 时间插入、删除和获取随机元素
 * @Date: 2024-05-17 15:25
 */

package main

import "math/rand"

type RandomizedSet struct {
	nums []int       // 可变长数组
	m    map[int]int // key: num -> value: index
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	index := len(this.nums)
	this.m[val] = index
	this.nums = append(this.nums, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.m[val]
	if !ok {
		return false
	}
	// 将最后一位数据覆盖当前数据
	last := len(this.nums) - 1
	this.nums[index] = this.nums[last]
	this.m[this.nums[index]] = index
	this.nums = this.nums[:last]
	delete(this.m, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	c := RandomizedSetConstructor()

	//c.Insert(0)
	//c.Insert(1)
	//c.Remove(0)
	//c.Insert(2)
	//c.Remove(1)
	//println(c.GetRandom())	// 2

	c.Remove(0)
	c.Remove(0)
	c.Insert(0)
	println(c.GetRandom()) //0
	c.Remove(0)
	c.Insert(0)

}
