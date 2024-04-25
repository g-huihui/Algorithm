/**
 * @Author: Gong Yanhui
 * @Description: 2739. 总行驶距离
 * @Date: 2024-04-25 19:17
 */

package main

func distanceTraveled(mainTank int, additionalTank int) int {
	var res int
	for mainTank != 0 {
		if mainTank >= 5 {
			mainTank -= 5
			res += 5
			if additionalTank > 0 {
				additionalTank -= 1
				mainTank += 1
			}
		} else {
			res += mainTank
			mainTank = 0
		}
	}

	return res * 10
}

func main() {
	println(distanceTraveled(5, 10)) // 60
	println(distanceTraveled(1, 2))  // 10
}
