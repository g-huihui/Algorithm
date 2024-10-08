/**
 * @Author: Gong Yanhui
 * @Description: 1436. 旅行终点站
 * @Date: 2024-10-08 10:23
 */

package main

func destCity(paths [][]string) string {
	var begin = make(map[string]struct{})
	var end = make(map[string]struct{})
	for _, path := range paths {
		begin[path[0]] = struct{}{}
		end[path[1]] = struct{}{}
	}
	for k := range end {
		if _, ok := begin[k]; !ok {
			return k
		}
	}

	return ""
}

func main() {
	println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}})) // Sao Paulo
	println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))                                  // A
	println(destCity([][]string{{"A", "Z"}}))                                                          // Z
}
