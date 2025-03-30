/**
 * @Author: Gong Yanhui
 * @Description: LCR 135. 报数
 * @Date: 2025-03-30 15:30
 */

package LCR

func countNumbers(cnt int) []int {
	var result []int

	var size = 1
	for i := 0; i < cnt; i++ {
		size *= 10
	}

	for i := 1; i < size; i++ {
		result = append(result, i)
	}

	return result
}
