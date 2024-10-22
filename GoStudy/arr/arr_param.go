/**
 * @Author: Gong Yanhui
 * @Description: arr作为参数传递
 * @Date: 2024-10-22 10:25
 */

package arr

import "fmt"

func ArrParam() {
	var arr = [3]int{1, 2, 3}
	changeArr(arr)
	fmt.Println("原数组: ", arr) // [1 2 3]

	arr = [3]int{1, 2, 3}
	changeArrPointer(&arr)
	fmt.Println("原数组: ", arr) // [100 2 3]

	// Array 作为参数传递 传递的是值 不会改变原数组
}

func changeArr(arr [3]int) {
	arr[0] = 100
}

func changeArrPointer(arr *[3]int) {
	(*arr)[0] = 100
}

// =================================================================================================

func SliceParam() {
	var arr = []int{1, 2, 3}
	changeSlice(arr)
	fmt.Println("原切片: ", arr) // [100 2 3]

	arr = []int{1, 2, 3}
	changeSlicePointer(&arr)
	fmt.Println("原切片: ", arr) // [100 2 3 4]

	// Slice 作为参数传递 传递的是引用 会改变原数组
	// 扩容情况下 只有指针传递才会改变原数组
}

func changeSlice(arr []int) {
	arr[0] = 100
	arr = append(arr, 4)
}

func changeSlicePointer(arr *[]int) {
	(*arr)[0] = 100
	*arr = append(*arr, 4)
}

// =================================================================================================

type ArrStruct struct {
	arr   [3]int
	slice []int
}

func StructParam() {
	var arr = ArrStruct{
		arr:   [3]int{1, 2, 3},
		slice: []int{1, 2, 3},
	}
	changeStruct(arr)
	fmt.Println("原结构体: ", arr) // {[1 2 3] [100 2 3]}

	arr = ArrStruct{
		arr:   [3]int{1, 2, 3},
		slice: []int{1, 2, 3},
	}
	changeStructPointer(&arr)
	fmt.Println("原结构体: ", arr) // {[100 2 3] [100 2 3 4]}
}

func changeStruct(arr ArrStruct) {
	arr.arr[0] = 100
	arr.slice[0] = 100
	arr.slice = append(arr.slice, 4)
}

func changeStructPointer(arr *ArrStruct) {
	arr.arr[0] = 100
	arr.slice[0] = 100
	arr.slice = append(arr.slice, 4)
}
