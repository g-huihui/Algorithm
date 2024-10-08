/**
 * @Author: Gong Yanhui
 * @Description: 数组拷贝
 * @Date: 2024-10-08 10:48
 */

package arr

import "fmt"

// ArrCopy 数组拷贝
func ArrCopy() {

	// 初始数组
	var arr = []int{1, 2, 3}
	// 拷贝数组
	var arrCopy1 = make([]int, len(arr)) // 必须初始化数组长度 不然copy()会panic
	copy(arrCopy1, arr)                  // 值拷贝

	arr[0] = 100          // 修改原数组
	fmt.Println(arrCopy1) // [1 2 3] 拷贝数组不受影响

	arrCopy1[0] = 200 // 修改拷贝数组
	fmt.Println(arr)  // [100 2 3] 原数组不受影响

	fmt.Println("--------------------------")

	// 初始数组
	arr = []int{1, 2, 3}
	// 拷贝数组
	var arrCopy2 []int
	arrCopy2 = append(arrCopy2, arr...) // 值拷贝

	arr[0] = 100          // 修改原数组
	fmt.Println(arrCopy2) // [1 2 3] 拷贝数组不受影响

	arrCopy2[0] = 200 // 修改拷贝数组
	fmt.Println(arr)  // [100 2 3] 原数组不受影响

	fmt.Println("--------------------------")

	// 初始数组
	arr = []int{1, 2, 3}
	// 拷贝数组
	var arrCopy3 = arr // 引用拷贝

	arr[0] = 100          // 修改原数组
	fmt.Println(arrCopy3) // [100 2 3] 拷贝数组受影响

	arrCopy3[0] = 200 // 修改拷贝数组
	fmt.Println(arr)  // [200 2 3] 原数组受影响
}
