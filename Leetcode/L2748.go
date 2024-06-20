/**
 * @Author: Gong Yanhui
 * @Description: 2748. 美丽下标对的数目
 * @Date: 2024-06-20 15:12
 */

package main

func countBeautifulPairs(nums []int) int {

	var ant int
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ant += cnt[y]
			}
		}
		for x >= 10 { // 这里需要 O(log x) 的时间
			x /= 10
		}
		cnt[x]++ // 统计最高位的出现次数
	}
	return ant
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
