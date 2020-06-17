package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image[sr][sc] == newColor {
		return image
	}

	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	r, c := len(image), len(image[0])

	if sr-1 > -1 && oldColor == image[sr-1][sc] {
		floodFill(image, sr-1, sc, newColor)
	}
	if sr+1 < r && oldColor == image[sr+1][sc] {
		floodFill(image, sr+1, sc, newColor)
	}
	if sc-1 > -1 && oldColor == image[sr][sc-1] {
		floodFill(image, sr, sc-1, newColor)
	}
	if sc+1 < c && oldColor == image[sr][sc+1] {
		floodFill(image, sr, sc+1, newColor)
	}

	return image
}
