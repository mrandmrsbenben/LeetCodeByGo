package main

import (
	"fmt"
	"math"
)

func main() {
	// common.MakeTestCases("largestTriangleArea")
	points1 := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Printf("Expect1: 2\n")
	fmt.Printf("Output1: %v\n", largestTriangleArea(points1))
	points2 := [][]int{{4, 6}, {6, 5}, {3, 1}}
	fmt.Printf("Expect2: 5.5\n")
	fmt.Printf("Output2: %v\n", largestTriangleArea(points2))
	points3 := [][]int{{8, 3}, {5, 6}, {3, 5}}
	fmt.Printf("Expect3: 4.5\n")
	fmt.Printf("Output3: %v\n", largestTriangleArea(points3))
	points4 := [][]int{{10, 7}, {0, 4}, {5, 7}}
	fmt.Printf("Expect4: 7.5\n")
	fmt.Printf("Output4: %v\n", largestTriangleArea(points4))
	points5 := [][]int{{9, 0}, {0, 2}, {3, 1}, {10, 8}}
	fmt.Printf("Expect5: 37.0\n")
	fmt.Printf("Output5: %v\n", largestTriangleArea(points5))
	points6 := [][]int{{2, 5}, {7, 7}, {10, 8}, {10, 10}, {1, 1}}
	fmt.Printf("Expect5: 14.5\n")
	fmt.Printf("Output5: %v\n", largestTriangleArea(points6))

}
func largestTriangleArea(points [][]int) float64 {
	f := func(p [][]int) float64 {
		var pe, a, b, c float64
		a = math.Sqrt(math.Pow(float64(p[0][0]-p[1][0]), 2) + math.Pow(float64(p[0][1]-p[1][1]), 2))
		b = math.Sqrt(math.Pow(float64(p[0][0]-p[2][0]), 2) + math.Pow(float64(p[0][1]-p[2][1]), 2))
		c = math.Sqrt(math.Pow(float64(p[1][0]-p[2][0]), 2) + math.Pow(float64(p[1][1]-p[2][1]), 2))
		pe = (a + b + c) / 2
		return math.Sqrt(pe * (pe - a) * (pe - b) * (pe - c))
	}
	var area, a float64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				a = f([][]int{points[i], points[j], points[k]})
				if area < a {
					area = a
				}
			}
		}
		if i == len(points)-3 {
			break
		}
	}
	return area
}
func largestTriangleArea0(points [][]int) float64 {
	f := func(p [][]int) float64 {
		var maxX, maxY, minX, minY, midX, midY int
		minX = p[0][0]
		minY = p[0][1]
		for _, v := range p {
			if v[0] < minX {
				minX = v[0]
			}
			if v[1] < minY {
				minY = v[1]
			}
			if v[0] > maxX {
				maxX = v[0]
			}
			if v[1] > maxY {
				maxY = v[1]
			}
		}
		area := float64((maxX - minX) * (maxY - minY))
		area = area - math.Abs(float64(p[0][0]-p[1][0]))*math.Abs(float64(p[0][1]-p[1][1]))/2
		area = area - math.Abs(float64(p[0][0]-p[2][0]))*math.Abs(float64(p[0][1]-p[2][1]))/2
		area = area - math.Abs(float64(p[1][0]-p[2][0]))*math.Abs(float64(p[1][1]-p[2][1]))/2
		fmt.Println(p, midX, midY, area)
		return area
	}
	var area, a float64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				a = f([][]int{points[i], points[j], points[k]})
				if area < a {
					area = a
				}
			}
		}
		if i == len(points)-3 {
			break
		}
	}
	return area
}
