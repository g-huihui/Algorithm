/**
 * @Author: Gong Yanhui
 * @Description: 303. 区域和检查 - 数组不可变
 * @Date: 2024-03-18 15:11
 */

package main

type NumArray struct {
	Nums    []int
	SumNums []int
}

func NumArrayConstructor(nums []int) NumArray {
	var res NumArray
	res.Nums = make([]int, len(nums))
	res.SumNums = make([]int, len(nums))
	for i, _ := range nums {
		res.Nums[i] = nums[i]
		if i == 0 {
			res.SumNums[i] = nums[i]
		} else {
			res.SumNums[i] = res.SumNums[i-1] + nums[i]
		}
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.SumNums[j] - this.SumNums[i] + this.Nums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
