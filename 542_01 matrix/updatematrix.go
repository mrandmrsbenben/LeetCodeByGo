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
func updateMatrix(matrix [][]int) [][]int {
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
