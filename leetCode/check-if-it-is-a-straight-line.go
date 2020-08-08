package main

func checkStraightLine(coordinates [][]int) bool {

	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]

	for i := range coordinates[:len(coordinates)-1] {

		a1, b1 := coordinates[i][0], coordinates[i][1]
		a2, b2 := coordinates[i+1][0], coordinates[i+1][1]

		if (b2-b1)*(x2-x1) != (y2-y1)*(a2-a1) {
			return false
		}

	}

	return true
}
