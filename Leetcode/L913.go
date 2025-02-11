/**
 * @Author: Gong Yanhui
 * @Description: 913. 猫和老鼠
 * @Date: 2025-02-11 10:43
 */

package main

func main() {
	// graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]	输出0
	println(catMouseGame([][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}))

	// graph = [[1,3],[0],[3],[0,2]]	输出1
	println(catMouseGame([][]int{{1, 3}, {0}, {3}, {0, 2}}))
}

const (
	mouseTurn = 0
	catTurn   = 1

	draw     = 0
	mouseWin = 1
	catWin   = 2
)

/**
方法一：拓扑排序 思路和算法
自顶向下的动态规划由于判定平局的标准和轮数有关，因此时间复杂度较高。为了降低时间复杂度，需要使用自底向上的方法实现，消除结果和轮数之间的关系。
使用自底向上的方法实现时，游戏中的状态由老鼠的位置、猫的位置和轮到移动的一方三个因素确定。初始时，只有边界情况的胜负结果已知，其余所有状态的结果都初始化为平局。边界情况为直接确定胜负的情况，包括两类情况：老鼠躲入洞里，无论猫位于哪个节点，都是老鼠获胜；猫和老鼠占据相同的节点，无论占据哪个节点，都是猫获胜。
从边界情况出发遍历其他情况。对于当前状态，可以得到老鼠的位置、猫的位置和轮到移动的一方，根据当前状态可知上一轮的所有可能状态，其中上一轮的移动方和当前的移动方相反，上一轮的移动方在上一轮状态和当前状态所在的节点不同。假设当前状态是老鼠所在节点是 mouse，猫所在节点是 cat，则根据当前的移动方，可以得到上一轮的所有可能状态：
如果当前的移动方是老鼠，则上一轮的移动方是猫，上一轮状态中老鼠所在节点是 mouse，猫所在节点可能是 graph[cat] 中的任意一个节点（除了节点 0）；
如果当前的移动方是猫，则上一轮的移动方是老鼠，上一轮状态中老鼠所在节点可能是 graph[mouse] 中的任意一个节点，猫所在节点是 cat。
对于上一轮的每一种可能的状态，如果该状态的结果已知不是平局，则不需要重复计算该状态的结果，只有对结果是平局的状态，才需要计算该状态的结果。对于上一轮的移动方，只有当可以确定上一轮状态是必胜状态或者必败状态时，才更新上一轮状态的结果。
如果上一轮的移动方和当前状态的结果的获胜方相同，由于当前状态为上一轮的移动方的必胜状态，因此上一轮的移动方一定可以移动到当前状态而获胜，上一轮状态为上一轮的移动方的必胜状态。
如果上一轮的移动方和当前状态的结果的获胜方不同，则上一轮的移动方需要尝试其他可能的移动，可能有以下三种情况：
如果存在一种移动可以到达上一轮的移动方的必胜状态，则上一轮状态为上一轮的移动方的必胜状态；
如果所有的移动都到达上一轮的移动方的必败状态，则上一轮状态为上一轮的移动方的必败状态；
如果所有的移动都不能到达上一轮的移动方的必胜状态，但是存在一种移动可以到达上一轮的移动方的必和状态，则上一轮状态为上一轮的移动方的必和状态。
其中，对于必败状态与必和状态的判断依据为上一轮的移动方可能的移动是都到达必败状态还是可以到达必和状态。为了实现必败状态与必和状态的判断，需要记录每个状态的度，初始时每个状态的度为当前玩家在当前位置可以移动到的节点数。对于老鼠而言，初始的度为老鼠所在的节点的相邻节点数；对于猫而言，初始的度为猫所在的节点的相邻且非节点 0 的节点数。
遍历过程中，从当前状态出发遍历上一轮的所有可能状态，如果上一轮状态的结果是平局且上一轮的移动方和当前状态的结果的获胜方不同，则将上一轮状态的度减 1。如果上一轮状态的度减少到 0，则从上一轮状态出发到达的所有状态都是上一轮的移动方的必败状态，因此上一轮状态也是上一轮的移动方的必败状态。
在确定上一轮状态的结果（必胜或必败）之后，即可从上一轮状态出发，遍历其他结果是平局的状态。当没有更多的状态可以确定胜负结果时，遍历结束，此时即可得到初始状态的结果。
细心的读者可以发现，上述遍历的过程其实是拓扑排序。

证明
必胜状态和必败状态都符合博弈中的最优策略，需要证明的是必和状态的正确性。
遍历结束之后，如果一个状态的结果是平局，则该状态满足以下两个条件：
从该状态出发，任何移动都无法到达该状态的移动方的必胜状态；
从该状态出发，存在一种移动可以到达必和状态。
对于标记结果是平局的状态，如果其实际结果是该状态的移动方必胜，则一定存在一个下一轮状态，为当前状态的移动方的必胜状态，在根据下一轮状态的结果标记当前状态的结果时会将当前状态标记为当前状态的移动方的必胜状态，和标记结果是平局矛盾。
对于标记结果是平局的状态，如果其实际结果是该状态的移动方必败，则所有的下一轮状态都为当前状态的移动方的必败状态，在根据下一轮状态的结果标记当前状态的结果时会将当前状态标记为当前状态的移动方的必败状态，和标记结果是平局矛盾。
因此，如果标记的状态是必和状态，则实际结果一定是必和状态。
*/

// 官方解法2
func catMouseGame(graph [][]int) int {
	n := len(graph)
	degrees := make([][][2]int, n)
	results := make([][][2]int, n)
	for i := range degrees {
		degrees[i] = make([][2]int, n)
		results[i] = make([][2]int, n)
	}
	for i, to := range graph {
		for j := 1; j < n; j++ {
			degrees[i][j][mouseTurn] = len(to)
			degrees[i][j][catTurn] = len(graph[j])
		}
	}
	for _, y := range graph[0] {
		for i := range degrees {
			degrees[i][y][catTurn]--
		}
	}

	type state struct{ mouse, cat, turn int }
	q := []state{}
	for j := 1; j < n; j++ {
		results[0][j][mouseTurn] = mouseWin
		results[0][j][catTurn] = mouseWin
		q = append(q, state{0, j, mouseTurn}, state{0, j, catTurn})
	}
	for i := 1; i < n; i++ {
		results[i][i][mouseTurn] = catWin
		results[i][i][catTurn] = catWin
		q = append(q, state{i, i, mouseTurn}, state{i, i, catTurn})
	}

	getPrevStates := func(s state) (prevStates []state) {
		if s.turn == mouseTurn {
			for _, prev := range graph[s.cat] {
				if prev != 0 {
					prevStates = append(prevStates, state{s.mouse, prev, catTurn})
				}
			}
		} else {
			for _, prev := range graph[s.mouse] {
				prevStates = append(prevStates, state{prev, s.cat, mouseTurn})
			}
		}
		return
	}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		result := results[s.mouse][s.cat][s.turn]
		for _, p := range getPrevStates(s) {
			prevMouse, prevCat, prevTurn := p.mouse, p.cat, p.turn
			if results[prevMouse][prevCat][prevTurn] == draw {
				canWin := result == mouseWin && prevTurn == mouseTurn || result == catWin && prevTurn == catTurn
				if canWin {
					results[prevMouse][prevCat][prevTurn] = result
					q = append(q, p)
				} else {
					degrees[prevMouse][prevCat][prevTurn]--
					if degrees[prevMouse][prevCat][prevTurn] == 0 {
						if prevTurn == mouseTurn {
							results[prevMouse][prevCat][prevTurn] = catWin
						} else {
							results[prevMouse][prevCat][prevTurn] = mouseWin
						}
						q = append(q, p)
					}
				}
			}
		}
	}
	return results[1][2][mouseTurn]
}

// 【宫水三叶】动态规划运用题 没理解透 通过JAVA改写Golang 计算出的结果不对
func catMouseGame2(graph [][]int) int {
	//var N = 55
	var n = len(graph)
	//var dp = make([][][]int, N)
	var dp = make([][][]int, 2*n*n)

	// 初始化dp递归数组
	for k := 0; k < 2*n*n; k++ {
		dp[k] = make([][]int, n)
		for i := 0; i < n; i++ {
			dp[k][i] = make([]int, n)
			for j := 0; j < n; j++ {
				dp[k][i][j] = -1
			}
		}
	}

	// 递归函数 0:draw 1:mouse 2:cat
	var dfs func(int, int, int) int
	dfs = func(k, a, b int) int {
		var ans = dp[k][a][b]
		if a == 0 {
			ans = 1
		} else if a == b {
			ans = 2
		} else if k >= 2*n*n {
			ans = 0
		} else if ans == -1 {
			if k%2 == 0 { // mouse回合
				var win, dra bool
				for _, ne := range graph[a] {
					var t = dfs(k+1, ne, b)
					if t == 1 {
						win = true
					} else if t == 0 {
						dra = true
					}
					if win {
						break
					}
				}
				if win {
					ans = 1
				} else if dra {
					ans = 0
				} else {
					ans = 2
				}
			} else { // cat回合
				var win, dra bool
				for _, ne := range graph[b] {
					if ne == 0 {
						continue
					}
					var t = dfs(k+1, a, ne)
					if t == 2 {
						win = true
					} else if t == 0 {
						dra = true
					}
					if win {
						break
					}
				}
				if win {
					ans = 2
				} else if dra {
					ans = 0
				} else {
					ans = 1
				}
			}
		}
		dp[k][a][b] = ans
		return ans
	}

	return dfs(1, 2, 0)
}

//const (
//	draw     = 0
//	mouseWin = 1
//	catWin   = 2
//)

// 官方解法1 超出时间限制
func catMouseGame1(graph [][]int) int {
	n := len(graph)
	// 初始化动态规划dp递归数组 dp[mouse][cat][turns] mouse所在节点 cat所在节点 turns进行的轮数
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n*(n-1)*2) // n*(n-1)*2 回合最大次数 mouse n个可能节点 cat n-1个可能节点
			for k := range dp[i][j] {
				// 没有遍历的点赋值为-1
				dp[i][j][k] = -1
			}
		}
	}

	var getResult, getNextResult func(int, int, int) int
	getResult = func(mouse, cat, turns int) int {
		if turns == n*(n-1)*2 { // 已经达到最大回合次数 平局
			return draw
		}
		res := dp[mouse][cat][turns]
		if res != -1 { // 1 mouse赢 2 cat赢 0 平局
			return res
		}
		if mouse == 0 { // mouse到达0洞口
			res = mouseWin
		} else if cat == mouse { // cat遇到mouse
			res = catWin
		} else {
			res = getNextResult(mouse, cat, turns) // 回溯
		}
		dp[mouse][cat][turns] = res
		return res
	}
	getNextResult = func(mouse, cat, turns int) int {
		// turns为偶数为mouse的回合 奇数为cat的回合
		curMove := mouse
		if turns%2 == 1 {
			curMove = cat
		}
		defaultRes := mouseWin
		if curMove == mouse {
			defaultRes = catWin
		}
		res := defaultRes
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				// 去除cat不能移动到0位置(洞)
				continue
			}
			nextMouse, nextCat := mouse, cat
			// 分别根据mouse和cat的回合将其移动到下一个位置
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}
			nextRes := getResult(nextMouse, nextCat, turns+1)
			if nextRes != defaultRes {
				res = nextRes
				if res != draw {
					break
				}
			}
		}
		return res
	}
	return getResult(1, 2, 0)
}
