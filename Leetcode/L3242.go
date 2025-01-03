/**
 * @Author: Gong Yanhui
 * @Description: 3242. 设计相邻元素求和服务
 * @Date: 2024-11-10 10:42
 */

package main

type Index struct {
	X int
	Y int
}

type NeighborSum struct {
	Grid     [][]int
	IndexMap map[int]Index
}

func NeighborSumConstructor(grid [][]int) NeighborSum {
	indexMap := make(map[int]Index)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			indexMap[grid[i][j]] = Index{i, j}
		}
	}
	return NeighborSum{
		Grid:     grid,
		IndexMap: indexMap,
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	index, ok := this.IndexMap[value]
	if !ok {
		return -1
	}

	var sum int
	if index.X > 0 {
		sum += this.Grid[index.X-1][index.Y]
	}
	if index.X < len(this.Grid)-1 {
		sum += this.Grid[index.X+1][index.Y]
	}
	if index.Y > 0 {
		sum += this.Grid[index.X][index.Y-1]
	}
	if index.Y < len(this.Grid[0])-1 {
		sum += this.Grid[index.X][index.Y+1]
	}

	return sum
}

func (this *NeighborSum) DiagonalSum(value int) int {
	index, ok := this.IndexMap[value]
	if !ok {
		return -1
	}

	var sum int
	if index.X > 0 && index.Y > 0 {
		sum += this.Grid[index.X-1][index.Y-1]
	}
	if index.X < len(this.Grid)-1 && index.Y < len(this.Grid[0])-1 {
		sum += this.Grid[index.X+1][index.Y+1]
	}
	if index.X > 0 && index.Y < len(this.Grid[0])-1 {
		sum += this.Grid[index.X-1][index.Y+1]
	}
	if index.X < len(this.Grid)-1 && index.Y > 0 {
		sum += this.Grid[index.X+1][index.Y-1]
	}

	return sum
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
