/**
 * @Author: Gong Yanhui
 * @Description: 12. 整数转罗马数字
 * @Date: 2024-06-05 11:56
 */

package main

func intToRoman(num int) string {
	var res string
	var m = map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, key := range keys {
		if (num / key) != 0 {
			count := num / key
			for i := 0; i < count; i++ {
				res += m[key]
			}
			num %= key
		}
	}

	return res
}

func main() {
	println(intToRoman(3749) == "MMMDCCXLIX")
	println(intToRoman(58) == "LVIII")
	println(intToRoman(1994) == "MCMXCIV")
}
