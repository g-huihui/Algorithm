/**
 * @Author: Gong Yanhui
 * @Description: 690. 员工的重要性
 * @Date: 2024-08-26 20:35
 */

package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	employeeMap := make(map[int]*Employee)
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	var dfs func(id int) int
	dfs = func(id int) int {
		employee := employeeMap[id]
		total := employee.Importance
		for _, subId := range employee.Subordinates {
			total += dfs(subId)
		}
		return total
	}

	return dfs(id)
}
