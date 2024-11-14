/**
 * @Author: Gong Yanhui
 * @Description: 3261. 统计满足 K 约束的子字符串数量 II
 * @Date: 2024-11-13 20:56
 */

package main

import "fmt"

// OOM
func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	var size = len(s)
	var dp = make([][]int64, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int64, size)
	}

	for begin := 0; begin < len(s); begin++ {
		var m = make(map[byte]int, 2)
		var count int64
		for end := begin; end < len(s); end++ {
			m[s[end]]++
			if m['1'] <= k || m['0'] <= k {
				count++
			}
			dp[begin][end] = count
		}
	}

	var resLen = len(queries)
	var res = make([]int64, 0, resLen)
	for i := 0; i < resLen; i++ {
		var sum int64
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			sum += dp[j][queries[i][1]]
		}
		res = append(res, sum)
	}

	return res
}

func main() {
	fmt.Println(countKConstraintSubstrings3("0001111", 2, [][]int{{0, 6}}))                // {26}
	fmt.Println(countKConstraintSubstrings3("010101", 1, [][]int{{0, 5}, {1, 4}, {2, 3}})) // {15, 9, 3}
}

/**
[0001111] 2

[1] [0] [0] [0] [0] [0] [0]
[2] [1] [0] [0] [0] [0] [0]
[3] [2] [1] [0] [0] [0] [0]
[4] [3] [2] [1] [0] [0] [0]
[5] [4] [3] [2] [1] [0] [0]
[5] [5] [4] [3] [2] [1] [0]
[5] [6] [5] [4] [3] [2] [1]	26

[0,6] => 26
*/

/**
[010101] 1

[1] [0] [0] [0] [0] [0]
[2] [1] [0] [0] [0] [0]
[3] [2] [1] [0] [0] [0]
[3] [3] [2] [1] [0] [0]
[3] [3] [3] [2] [1] [0]
[3] [3] [3] [3] [2] [1]	15

[0,5] => 15
[1,4] => 9
[2,3] => 3
[n,m] => 第m行 n-m的和
*/

// 优化一版 将数组折半 内存减半 依旧OOM
func countKConstraintSubstrings3(s string, k int, queries [][]int) []int64 {
	var size = len(s)
	var dp = make([][]int64, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int64, size-i)
	}

	for begin := 0; begin < len(s); begin++ {
		var m = make(map[byte]int, 2)
		var count int64
		for end := begin; end < len(s); end++ {
			m[s[end]]++
			if m['1'] <= k || m['0'] <= k {
				count++
			}
			dp[begin][end-begin] = count
		}
	}

	var resLen = len(queries)
	var res = make([]int64, 0, resLen)
	for i := 0; i < resLen; i++ {
		var sum int64
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			sum += dp[j][queries[i][1]-j]
		}
		res = append(res, sum)
	}

	return res
}

/**
实际存储格式:
[1] [2] [3] [3] [3] [3]
    [1] [2] [3] [3] [3]
        [1] [2] [3] [3]
            [1] [2] [3]
				[1] [2]
					[1]

[1]
[2] [1]
[3] [2] [1]
[3] [3] [2] [1]
[3] [3] [3] [2] [1]
[3] [3] [3] [3] [2] [1]
*/
