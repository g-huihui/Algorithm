/**
 * @Author: Gong Yanhui
 * @Description: 贷款利息计算
 * @Date: 2024-02-20 11:17
 */

package something

import "fmt"

func Calculate() {
	totalPrincipal, ms := getMoney()
	var (
		total     int
		principal int
		interest  int
	)

	for _, m := range ms {
		if !m.check() {
			panic(fmt.Sprintf("数据错误:%+v", m))
		}
		total += m.Total
		principal += m.Principal
		interest += m.Interest
	}

	if principal != totalPrincipal {
		panic(fmt.Sprintf("数据错误: principal != totalPrincipal principal:%f,totalPrincipal:%f", principal, totalPrincipal))
	}

	if total != totalPrincipal+interest {
		panic(fmt.Sprintf("数据错误: total != totalPrincipal+interest total:%f,totalPrincipal:%f,interest:%f", total, totalPrincipal, interest))
	}
	fmt.Printf("total:%d,principal:%d,interest:%d", total, principal, interest)
}

func getMoney() (int, []Money) {
	// 2023
	var totalPrincipal int = 2619103
	var ms = []Money{
		{Total: 1707023, Principal: 369398, Interest: 1337625}, // 7月
		{Total: 1261148, Principal: 370660, Interest: 890488},  // 8月
		{Total: 1261148, Principal: 371927, Interest: 889221},  // 9月
		{Total: 1261148, Principal: 373197, Interest: 887951},  // 10月
		{Total: 1261148, Principal: 374472, Interest: 886676},  // 11月
		{Total: 1261148, Principal: 375752, Interest: 885396},  // 12月
		{Total: 1267809, Principal: 383697, Interest: 884112},  // 1月
	}

	return totalPrincipal, ms
}

type Money struct {
	Total     int
	Principal int
	Interest  int
}

func (m Money) check() bool {
	return m.Total == m.Principal+m.Interest
}
