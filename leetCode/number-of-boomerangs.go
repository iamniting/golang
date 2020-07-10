package main

func numberOfBoomerangs(points [][]int) int {
	var res int

	for _, p := range points {

		m := make(map[int]int, len(points))

		for _, q := range points {
			d := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])

			if v, ok := m[d]; ok {
				res += 2 * v
			}
			m[d]++
		}
	}

	return res
}
