package main

import "fmt"

func main() {
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	sr := 1
	sc := 1
	newColor := 1
	fmt.Printf("Output: %v\n", floodFill(image, sr, sc, newColor))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	image[sr][sc] = newColor
	if sr > 0 && image[sr-1][sc] == oldColor {
		image = floodFill(image, sr-1, sc, newColor)
	}
	if sr+1 < len(image) && image[sr+1][sc] == oldColor {
		image = floodFill(image, sr+1, sc, newColor)
	}
	if sc > 0 && image[sr][sc-1] == oldColor {
		image = floodFill(image, sr, sc-1, newColor)
	}
	if sc+1 < len(image[sr]) && image[sr][sc+1] == oldColor {
		image = floodFill(image, sr, sc+1, newColor)
	}
	return image
}
