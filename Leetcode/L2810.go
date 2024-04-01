/**
 * @Author: Gong Yanhui
 * @Description: 2810. 故障键盘
 * @Date: 2024-04-01 19:57
 */

package main

func finalString(s string) string {
	var res []byte
	var head = false
	for _, cha := range s {
		if cha != 'i' {
			if !head {
				res = append(res, byte(cha))
			} else {
				res = append([]byte{byte(cha)}, res...)
			}
		} else {
			head = !head
		}
	}
	if head {
		var left, right = 0, len(res) - 1
		for left < right {
			res[left], res[right] = res[right], res[left]
			left++
			right--
		}
	}
	return string(res)
}

func main() {
	if str := finalString("string"); str != "rtsng" {
		panic(str)
	}
	if str := finalString("poiinter"); str != "ponter" {
		panic(str)
	}
}
