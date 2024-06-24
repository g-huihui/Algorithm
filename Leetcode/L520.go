/**
 * @Author: Gong Yanhui
 * @Description: 520. 检测大写字母
 * @Date: 2024-06-24 21:02
 */

package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	var firstUpper, secondUpper bool
	if isUpper(word[0]) {
		firstUpper = true
	}
	if isUpper(word[1]) {
		secondUpper = true
	}
	if !firstUpper && secondUpper {
		return false
	}

	for i := 2; i < len(word); i++ {
		if isUpper(word[i]) != secondUpper {
			return false
		}
	}

	return true
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func main() {
	println(detectCapitalUse("USA"))  // true
	println(detectCapitalUse("FlaG")) // false
}
