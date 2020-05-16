// https://leetcode.com/problems/lucky-numbers-in-a-matrix
// Just sol to the problem, It does not include the I/O part

func luckyNumbers(matrix [][]int) []int {
	R, C, res := make([]int, len(matrix)), make([]int, len(matrix[0])), []int{}

	for i, row := range matrix {
		R[i] = math.MaxInt32
		for j, col := range row {
			R[i] = int(math.Min(float64(col), float64(R[i])))
			C[j] = int(math.Max(float64(col), float64(C[j])))
		}
	}

	for i, row := range matrix {
		for j, col := range row {
			if R[i] == col && C[j] == col {
				res = append(res, col)
			}
		}
	}
	return res
}

