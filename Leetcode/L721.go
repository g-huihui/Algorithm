/**
 * @Author: Gong Yanhui
 * @Description: 721. 账户合并
 * @Date: 2024-07-15 23:03
 */

package main

import "sort"

// 此方法没有解决多个交叉关联的问题 A和B没有关联 但C和AB都有关联 既ABC都有关联
func accountsMerge(accounts [][]string) [][]string {

	var nameMap = make(map[string][]int)          // name -> index
	var emailArr = make([]map[string]struct{}, 0) // index -> email

	for _, account := range accounts {
		name := account[0]
		indexs, ok := nameMap[name]
		if ok {
			index := isOneAccount(indexs, emailArr, account)
			if index != -1 {
				emailTmp := emailArr[index]
				for i := 1; i < len(account); i++ {
					emailTmp[account[i]] = struct{}{}
				}
				emailArr[index] = emailTmp
			} else {
				var emailTmp = make(map[string]struct{})
				for i := 1; i < len(account); i++ {
					emailTmp[account[i]] = struct{}{}
				}
				emailArr = append(emailArr, emailTmp)
				nameMap[name] = append(indexs, len(emailArr)-1)
			}
		} else {
			var emailTmp = make(map[string]struct{})
			for i := 1; i < len(account); i++ {
				emailTmp[account[i]] = struct{}{}
			}
			emailArr = append(emailArr, emailTmp)
			nameMap[name] = append([]int{}, len(emailArr)-1)
		}
	}

	var res = make([][]string, 0)
	for name, indexs := range nameMap {
		for _, index := range indexs {
			accs := emailArr[index]
			var tmp []string
			for acc, _ := range accs {
				tmp = append(tmp, acc)
			}
			// 按字符 ASCII 顺序排列
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] < tmp[j]
			})
			res = append(res, append([]string{name}, tmp...))
		}
	}

	return res
}

func isOneAccount(indexs []int, arr []map[string]struct{}, account []string) int {
	for _, index := range indexs {
		emailTmp := arr[index]
		for i := 1; i < len(account); i++ {
			_, ok := emailTmp[account[i]]
			if ok {
				return index
			}
		}
	}
	return -1
}

func main() {
	param1 := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"}}
	println(accountsMerge(param1))
}
