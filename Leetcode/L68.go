/**
 * @Author: Gong Yanhui
 * @Description: 68. 文本左右对齐
 * @Date: 2024-06-19 10:54
 */

package main

import "reflect"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var indexArr []int
	var curLen int
	for i := 0; i <= len(words); i++ {
		if i == len(words) && len(indexArr) > 0 {
			res = append(res, justify(words, indexArr, maxWidth, true))
			break
		}
		if curLen+len(words[i]) <= maxWidth {
			curLen += len(words[i]) + 1
			indexArr = append(indexArr, i)
		} else {
			res = append(res, justify(words, indexArr, maxWidth, false))
			indexArr = []int{i}
			curLen = len(words[i]) + 1
		}
	}

	return res
}

func justify(words []string, indexArr []int, maxWidth int, isLast bool) string {
	var res string
	if isLast {
		for i := 0; i < len(indexArr); i++ {
			res += words[indexArr[i]]
			if i < len(indexArr)-1 {
				res += " "
			}
		}
		count := maxWidth - len(res)
		for i := 0; i < count; i++ {
			res += " "
		}
		return res
	}

	if len(indexArr) == 1 {
		res = words[indexArr[0]]
		for i := 0; i < maxWidth-len(words[indexArr[0]]); i++ {
			res += " "
		}
		return res
	}

	var wordLen int
	for i := 0; i < len(indexArr); i++ {
		wordLen += len(words[indexArr[i]])
	}

	var spaceNum = len(indexArr) - 1
	var spaceLen = maxWidth - wordLen
	var lastSpace = spaceLen % spaceNum
	var avgSpace = spaceLen / spaceNum
	for i := 0; i < len(indexArr); i++ {
		res += words[indexArr[i]]
		if i != len(indexArr)-1 {
			for j := 0; j < avgSpace; j++ {
				res += " "
			}
			if lastSpace > 0 {
				res += " "
				lastSpace--
			}
		}
	}

	return res
}

func main() {
	res1 := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	except1 := []string{"This    is    an", "example  of text", "justification.  "}
	println(reflect.DeepEqual(res1, except1))

	res2 := fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	except2 := []string{"What   must   be", "acknowledgment  ", "shall be        "}
	println(testEq(res2, except2))

	res3 := fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to",
		"a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	except3 := []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
		"everything  else  we", "do                  "}
	println(testEq(res3, except3))
}

func testEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
