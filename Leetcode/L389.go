/**
 * @Author: Gong Yanhui
 * @Description: 389. 找不同
 * @Date: 2024-01-08 11:57
 */

package main

func findTheDifference(s string, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}
	return res
}
