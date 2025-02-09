/**
 * @Author: Gong Yanhui
 * @Description: 860. 柠檬水找零
 * @Date: 2025-02-09 20:13
 */

package main

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else { // bill == 20
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
