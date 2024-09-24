/**
 * @Author: Gong Yanhui
 * @Description: 打印方式
 * @Date: 2024-09-24 11:00
 */

package something

import "fmt"

func PrintColor() {
	fmt.Println("\033[35m" + "这是带颜色的输出" + "\033[0m")
}
