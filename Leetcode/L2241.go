/**
 * @Author: Gong Yanhui
 * @Description: 2241. 设计一个 ATM 机器
 * @Date: 2025-01-05 11:48
 */

package main

type ATM struct {
	size      int   // 记录总共有多少种不同面额
	bankNotes []int // 记录每种面额的数量
	bankValue []int // 记录每种面额的价值
}

func ATMConstructor() ATM {
	return ATM{
		size:      5,
		bankNotes: make([]int, 5),
		bankValue: []int{20, 50, 100, 200, 500},
	}
}

// Deposit [0,0,1,2,1]
func (this *ATM) Deposit(banknotesCount []int) {
	for i, n := range banknotesCount {
		this.bankNotes[i] += n
	}
}

func (this *ATM) Withdraw(amount int) []int {
	// 记录取走的每种金额的数量
	var res = make([]int, this.size)
	// 从大面额开始取钱
	for i := this.size - 1; i >= 0; i-- {
		// 计算需要的当前面额的数量
		var n = amount / this.bankValue[i]
		// 如果当前面额的数量大于存储的数量，取存储的数量
		if n > this.bankNotes[i] {
			n = this.bankNotes[i]
		}
		// 更新取走的数量
		res[i] = n
		// 更新取走的总金额
		amount -= n * this.bankValue[i]
	}
	// 如果取走的总金额不为0 则无法取款
	if amount != 0 {
		return []int{-1}
	}
	// 更新存储的数量
	for i := 0; i < this.size; i++ {
		this.bankNotes[i] -= res[i]
	}
	return res
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
