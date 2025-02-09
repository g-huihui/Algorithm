/**
 * @Author: Gong Yanhui
 * @Description: 1041. 困于环中的机器人
 * @Date: 2025-02-09 15:08
 */

package main

// 开始时 机器人位于远点并且方向向北
// 根据推算 只要执行完所有指令后 机器人不位于原点且方向朝北 则一定不会回到原点
func isRobotBounded(instructions string) bool {

	var direc = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 分别代表 向北、向东、向南、向西移动

	var x, y = 0, 0 // 机器人的坐标
	var idx = 0     // 机器人的方向

	// 执行指令
	for _, v := range instructions {
		if v == 'G' {
			x += direc[idx][0]
			y += direc[idx][1]
		} else if v == 'L' {
			idx = (idx + 3) % 4
		} else if v == 'R' {
			idx = (idx + 1) % 4
		}
	}

	return (x == 0 && y == 0) || idx != 0
}
