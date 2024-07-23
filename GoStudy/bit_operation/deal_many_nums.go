/**
 * @Author: Gong Yanhui
 * @Description: 通过二进制处理大量数据防止内存溢出 https://blog.csdn.net/u010983881/article/details/75097358
 * @Date: 2024-07-23 22:27
 */

package bit_operation

// DealManyNums 10亿数据 找出其中只出现一次的数 如果全部载入内存 10亿 * 4B = 4GB 内存 无法承受
// 通过位运算处理
func DealManyNums() {
	var nums = []int{1, 2, 2, 3, 4, 4, 5, 5, 5, 5, 6, 7, 8, 8}

	// 假设数据大小不超过32
	var tag uint32 = 0

	// 0 代表出现0次 1 代表出现1次
	for _, num := range nums {
		var tmp uint32 = 1 << num
		if tag&tmp == 0 {
			tag = tag | tmp
		} else {
			tag = tag ^ tmp
		}
	}

	var res []int
	for i := 0; i < 32; i++ {
		if tag&(1<<i) != 0 {
			res = append(res, i)
		}
	}

	println(res) // [1 3 6 7]
}
