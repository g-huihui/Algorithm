/**
 * @Author: Gong Yanhui
 * @Description: 2024. 考试的最大困扰度
 * @Date: 2024-09-02 20:12
 */

package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutive(answerKey, k, 'T'), maxConsecutive(answerKey, k, 'F'))
}

func maxConsecutive(key string, k int, ch rune) int {
	var leftIndex, maxCount = 0, -1
	var changeIndex = make([]int, 0, k)
	for index, v := range key {
		if v == ch {
			maxCount = max(maxCount, index-leftIndex+1)
			continue
		}
		if len(changeIndex) == k {
			leftIndex = changeIndex[0] + 1
			changeIndex = changeIndex[1:]
		}
		changeIndex = append(changeIndex, index)
		maxCount = max(maxCount, index-leftIndex+1)
	}
	return maxCount
}

func main() {
	println(maxConsecutiveAnswers("TTFF", 2))     // 4
	println(maxConsecutiveAnswers("TFFT", 1))     // 3
	println(maxConsecutiveAnswers("TTFTTFTT", 1)) // 5
}
