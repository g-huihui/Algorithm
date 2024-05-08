/**
 * @Author: Gong Yanhui
 * @Description: 2079. 给植物浇水
 * @Date: 2024-05-08 15:22
 */

package main

func wateringPlants(plants []int, capacity int) int {
	var pathCount = 0
	var water = capacity
	for index, val := range plants {
		if val > water {
			pathCount = pathCount + (index*2 + 1)
			water = capacity - val
		} else {
			pathCount += 1
			water -= val
		}
	}
	return pathCount
}

func main() {
	println(wateringPlants([]int{2, 2, 3, 3}, 5))          // 14
	println(wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4))    // 30
	println(wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8)) // 49
}
