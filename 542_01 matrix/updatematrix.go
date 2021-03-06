package main

import (
	"fmt"
	"strconv"
)

func main() {
	// matrix := [][]int{
	// 	{0, 0, 0},
	// 	{0, 1, 0},
	// 	{0, 0, 0}}
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1}}
	fmt.Println("Output:", updateMatrix(matrix))
}

//执行用时 :64 ms, 在所有 Go 提交中击败了91.25%的用户
//内存消耗 :7.3 MB, 在所有 Go 提交中击败了100.00%的用户
func updateMatrix(matrix [][]int) [][]int {
	data := [][2]int{}
	lr, lc := len(matrix), len(matrix[0])
	res := make([][]int, lr)
	for i := range matrix {
		res[i] = make([]int, lc)
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				data = append(data, [2]int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}

	var buf [][2]int
	distance := 1
	var x, y int
	for len(data) > 0 {
		buf = [][2]int{}
		for i := range data {
			x, y = data[i][1], data[i][0]
			// up
			if y > 0 && res[y-1][x] == -1 {
				res[y-1][x] = distance
				buf = append(buf, [2]int{y - 1, x})
			}
			// down
			if y < lr-1 && res[y+1][x] == -1 {
				res[y+1][x] = distance
				buf = append(buf, [2]int{y + 1, x})
			}
			// left
			if x > 0 && res[y][x-1] == -1 {
				res[y][x-1] = distance
				buf = append(buf, [2]int{y, x - 1})
			}
			// right
			if x < lc-1 && res[y][x+1] == -1 {
				res[y][x+1] = distance
				buf = append(buf, [2]int{y, x + 1})
			}
		}
		distance++
		data = make([][2]int, len(buf))
		copy(data, buf)
	}
	return res
}

//Point Array Point
type Point struct {
	x int
	y int
}

func (p *Point) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

//执行用时 :408 ms, 在所有 Go 提交中击败了73.17%的用户
//内存消耗 :138.1 MB, 在所有 Go 提交中击败了12.50%的用户
func updateMatrixV1(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	lr, lc := len(matrix), len(matrix[0])
	res := make([][]int, lr)
	f := func(r, c int) int {
		distance := 0
		pmap := make(map[string]int)
		p := make([]*Point, 1)
		p[0] = &Point{r, c}
		pmap[p[0].String()] = 1
		var newP *Point
		for len(p) > 0 {
			distance++
			l := len(p)
			for i := 0; i < l; i++ {
				if p[i].x+1 < lr {
					if matrix[p[i].x+1][p[i].y] == 0 {
						return distance
					}
					newP = &Point{p[i].x + 1, p[i].y}
					if pmap[newP.String()] == 0 {
						p = append(p, newP)
						pmap[newP.String()] = 1
					}
				}
				if p[i].x-1 >= 0 {
					if matrix[p[i].x-1][p[i].y] == 0 {
						return distance
					}
					newP = &Point{p[i].x - 1, p[i].y}
					if pmap[newP.String()] == 0 {
						p = append(p, newP)
						pmap[newP.String()] = 1
					}
				}
				if p[i].y+1 < lc {
					if matrix[p[i].x][p[i].y+1] == 0 {
						return distance
					}
					newP = &Point{p[i].x, p[i].y + 1}
					if pmap[newP.String()] == 0 {
						p = append(p, newP)
						pmap[newP.String()] = 1
					}
				}
				if p[i].y-1 >= 0 {
					if matrix[p[i].x][p[i].y-1] == 0 {
						return distance
					}
					newP = &Point{p[i].x, p[i].y - 1}
					if pmap[newP.String()] == 0 {
						p = append(p, newP)
						pmap[newP.String()] = 1
					}
				}
			}
			p = p[l:]
		}
		return distance
	}

	for i := range res {
		res[i] = make([]int, lc)
		for j := range res[i] {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			} else if i > 0 && res[i-1][j] == 0 {
				res[i][j] = 1
			} else if j > 0 && res[i][j-1] == 0 {
				res[i][j] = 1
			} else {
				res[i][j] = f(i, j)
			}
		}
	}
	return res
}
