package main

import "fmt"

func main() {
	instructions := "GG"
	fmt.Println("Output:", isRobotBounded(instructions))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有Go提交中击败了100.00%的用户
func isRobotBounded(instructions string) bool {
	d := 1 //directions :{'W', 'N', 'E', 'S'}
	x, y := 0, 0
	for _, i := range instructions {
		switch i {
		case 'L':
			if d == 0 {
				d = 3
			} else {
				d--
			}
		case 'R':
			if d == 3 {
				d = 0
			} else {
				d++
			}
		default:
			switch d {
			case 0:
				x--
			case 1:
				y++
			case 2:
				x++
			case 3:
				y--
			}
		}
	}
	if x == 0 && y == 0 {
		return true
	} else if d != 1 {
		return true
	}
	return false
}
