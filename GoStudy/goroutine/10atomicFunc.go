/**
 * @Author: Gong Yanhui
 * @Description: sync 包（三）：原子操作	https://laravelacademy.org/post/19993
 * @Date: 2022-11-22 20:58
 */

package main

import (
	"fmt"
	"sync/atomic"
)

// 原子操作的相关方法
func atomicAddFunc() {
	// 第一个参数是操作数对应的指针 第二个参数是加/减值 虽然都是以Add前缀开头 但是对于减法可以通过传递负数实现
	var i32 int32 = 0
	var addInt32 = atomic.AddInt32(&i32, 1)
	fmt.Println("AddInt32:", addInt32)

	var i64 int64 = 3
	var addInt64 = atomic.AddInt64(&i64, -1)
	fmt.Println("AddInt64:", addInt64)

	// 对于接下来三个函数 由于操作数类型是无符号的 所以无法显示传递负值来实现减法
	var iu32 uint32 = 2
	var addUint32 = atomic.AddUint32(&iu32, 1)
	fmt.Println("AddUint32:", addUint32)

	var iu64 uint64 = 3
	var addUint64 = atomic.AddUint64(&iu64, 1)
	fmt.Println("AddUint64:", addUint64)

	var iUintPtr uintptr = 4
	var addUintPtr = atomic.AddUintptr(&iUintPtr, 1)
	fmt.Println("AddUintptr:", addUintPtr)
}

// CAS方式原子操作相关方法
func compareAndSwapFunc() {
	// 第一个参数是操作数对应的指针 第二 三个参数是待比较和交换的旧值与新值 (addr, old, new) bool
	// 这些函数会在交换之前先判断addr地址中的值是否与old相等 如果不相等则返回false 否则将其替换成new
	var numInt32 int32 = 1
	atomic.CompareAndSwapInt32(&numInt32, numInt32, numInt32+1)
	fmt.Println("CompareAndSwapInt32:", numInt32)

	var numInt64 int64 = 2
	atomic.CompareAndSwapInt64(&numInt64, numInt64, numInt64+1)
	fmt.Println("CompareAndSwapInt64:", numInt64)

	//var numPointer unsafe.Pointer

	var numUInt32 uint32 = 3
	atomic.CompareAndSwapUint32(&numUInt32, numUInt32, numUInt32+1)
	fmt.Println("CompareAndSwapUint32:", numUInt32)

	var numUInt64 uint64 = 4
	atomic.CompareAndSwapUint64(&numUInt64, numUInt64, numUInt64+1)
	fmt.Println("CompareAndSwapUint64:", numUInt64)

	var numUIntPtr uintptr = 5
	atomic.CompareAndSwapUintptr(&numUIntPtr, numUIntPtr, numUIntPtr+1)
	fmt.Println("CompareAndSwapUintptr:", numUIntPtr)
}

// 加载相关的原子操作方法
func loadFunc() {
	// 这些操作函数仅传递一个参数 即待操作数对应的指针 并且有一个返回值 返回传入指针指向的值
	// 这里的「原子性」指的是当读取该指针指向的值时 CPU不会执行任何其它针对此值的读写操作
	var x int32 = 100
	y := atomic.LoadInt32(&x)
	fmt.Println("x, y:", x, y)

	// ...一共六种
}

// 存储相关的原子操作方法
func storeFunc() {
	// 第一个参数表示待操作变量对应的指针 第二个参数表示要存储到待操作变量的数值
	// 该操作可以看作是加载操作的逆向操作 一个用于读取 一个用于写入 通过上述原子函数存储数值的时候 不会出现存储流程进行到一半被中断的情况
	var x int32 = 100
	var y int32
	atomic.StoreInt32(&y, atomic.LoadInt32(&x))
	fmt.Println("x, y:", x, y)

	// ...一共六种
}

// 交换相关的原子操作方法
func swapFunc() {
	// 交换和比较并交换看起来有点类似 但是交换不关心待操作数的旧值 不管旧值和新值是否相等 都会通过新值替换旧值
	// 交换函数有一个返回值，会返回旧值
	var num int32 = 1
	oldNum := atomic.SwapInt32(&num, 100)
	fmt.Println("old =", oldNum, " new =", num)
}

// 原子变量
func atomicValue() {
	// 声明一个该类型的变量之后就可以直接使用了 这个类型使用起来很简单 它只有 Store 和 Load 两个指针方法 这两个方法都是原子操作
	// 还是有一些需要注意的地方 首先存储值不能是 nil 其次我们向原子类型存储的第一个值 决定了它今后能且只能存储该类型的值 如果违背这两条 编译时会抛出 panic
	var v atomic.Value
	v.Store(100)
	fmt.Println(v.Load())
}

func main() {
	atomicAddFunc()
	compareAndSwapFunc()
	loadFunc()
	storeFunc()
	swapFunc()
	atomicValue()
}
