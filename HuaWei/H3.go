/**
 * @Author: Gong Yanhui
 * @Description: 查找充电设备组合
 * @Date: 2024-12-21 14:37
 */

package main

/**
1. 题目描述
某个充电站 可提供n个充电设备 每个充电设备均有对应的輸出功率
任意个充电设备组合的输出功率总和 构成的功率集合P表示 最接近充电站最大输出功率p_max的元素

2. 输入描述
n		充电设备个数
arr		每个充电设备输出功率
p_max	充电站最大输出功率

3. 输入输出
input : 5 [50,20,20,60] 90
output: 90 (50+20+20)
input : 2 [50,40] 30
output: 0
*/

func H3(n int, arr []int, pMax int) int {

	var res int

	// 双指针扫描 获取最大值
	var begin, end int
	var total int
	for end < len(arr) || begin > end {
		if total+arr[end] <= pMax {
			total += arr[end]
			end++
			if total > res {
				res = total
			}
		} else {
			total -= arr[begin]
			begin++
		}
	}

	return res
}

func main() {
	println(H3(5, []int{50, 20, 20, 60}, 90)) // 90
	println(H3(2, []int{50, 40}, 30))         // 0

	println(H3(1, []int{50}, 30))        // 0
	println(H3(1, []int{50}, 50))        // 50
	println(H3(0, []int{}, 50))          // 0
	println(H3(4, []int{1, 2, 3, 4}, 6)) // 6
	println(H3(4, []int{1, 2, 3, 4}, 9)) // 9
}
