/**
 * @Author: Gong Yanhui
 * @Description: map相关问题
 * @Date: 2024-01-08 14:46
 */

package something

import (
	"fmt"
	"time"
)

// TestConcurrentMap 测试map并发触发panic问题 fatal error: concurrent map writes
func TestConcurrentMap() {
	m := make(map[int]int)
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			m[i] = i // 写
		}
		fmt.Println("write done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			_ = m[i] // 读
		}
		fmt.Println("read done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			m[i] = i // 写
		}
		fmt.Println("write done")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 1000; i++ {
			_ = m[i] // 读
		}
		fmt.Println("read done")
	}()

	time.Sleep(10 * time.Second)
}

func TestConcurrentMap2() {
	// 测试map只并发读 不会出现panic
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[30] = 30
	m[40] = 40
	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			for j := 0; j < 1000; j++ {
				_ = m[j] // 读
			}
			fmt.Println("read done")
		}()
	}

	time.Sleep(10 * time.Second)
}

func TestConcurrentMap3() {
	// 测试map只并发写 fatal error: concurrent map writes
	m := make(map[int]int)
	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			for j := 0; j < 1000; j++ {
				m[j] = j // 写
			}
			fmt.Println("write done")
		}()
	}

	time.Sleep(10 * time.Second)
}

func TestConcurrentMap4() {
	// 测试map只并发读 只有一个协程写 fatal error: concurrent map read and map write
	m := make(map[int]int)
	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			for j := 0; j < 1000; j++ {
				_ = m[j] // 读
			}
			fmt.Println("read done")
		}()
	}

	go func() {
		time.Sleep(1 * time.Second)
		for j := 0; j < 10000; j++ {
			m[j] = j // 写
		}
		fmt.Println("write done")
	}()

	time.Sleep(10 * time.Second)
}

// MapStruct 测试map中的struct赋值问题
func MapStruct() {
	type User struct {
		Name string
	}
	type Snapshot struct {
		Key   string
		Users []User
	}

	snapshots := make(map[string]Snapshot, 1)
	snapshots["test"] = Snapshot{
		Key:   "testKey",
		Users: make([]User, 0),
	}

	user := User{"testUser"}

	// 1. 不能直接修改map中的struct
	// snapshots["test"].Users = append(snapshots["test"].Users, user)
	// error: Cannot assign to snapshots["test"].Users

	// 2. 可以先取出来修改再放回去
	tmp := snapshots["test"]
	tmp.Users = append(tmp.Users, user)
	snapshots["test"] = tmp

	// 3. 尝试定义指针类型
	snapshotsPointer := make(map[string]*Snapshot, 1)
	snapshotsPointer["test"] = &Snapshot{
		Key:   "testKey",
		Users: make([]User, 0),
	}
	snapshotsPointer["test"].Users = append(snapshotsPointer["test"].Users, user)

	println(len(snapshots))
	println(snapshots["test"].Users[0].Name)
	println(len(snapshotsPointer))
	println(snapshotsPointer["test"].Users[0].Name)
}

func MapLength() {
	a := make([]int, 3)
	fmt.Println(len(a)) // 3

	b := make(map[string]string, 3)
	fmt.Println(len(b)) // 0

	// make(type, len, cap) cap == 0
}

func TestMapIsPointer() {
	m1 := make(map[int]struct{})
	testMapAppend(m1)
	fmt.Println(len(m1)) // 2

	m2 := make(map[int]struct{})
	testMapPointAppend(&m2)
	fmt.Println(len(m2)) // 4

	// 结论: map是引用类型 传递指针和传递值的没有区别
}

func testMapAppend(m map[int]struct{}) {
	m[1] = struct{}{}
	m[4] = struct{}{}
}

func testMapPointAppend(m *map[int]struct{}) {
	(*m)[1] = struct{}{}
	(*m)[8] = struct{}{}
	(*m)[9] = struct{}{}
	(*m)[11] = struct{}{}
}
