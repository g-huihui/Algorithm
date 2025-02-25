/**
 * @Author: Gong Yanhui
 * @Description: 2502. 设计内存分配器
 * @Date: 2025-02-25 22:20
 */

package main

type Allocator struct {
	// 内存大小
	size int
	// 内存分配表
	memory []int // 0表示未分配
}

func AllocatorConstructor(n int) Allocator {
	return Allocator{
		size:   n,
		memory: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	var free int // 记录连续空闲内存的大小
	for i := 0; i < this.size; i++ {
		if this.memory[i] == 0 {
			free++
			if free == size { // 达到要求的内存大小
				// 将这块内存空间赋值为mID
				for j := 0; j < size; j++ {
					this.memory[i-j] = mID
				}
				return i - size + 1
			}
		} else {
			free = 0
		}
	}

	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	var count = 0
	// 循环遍历内存空间为mID的位置将其释放(值为0)
	for i := 0; i < this.size; i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			count++
		}
	}

	return count
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
