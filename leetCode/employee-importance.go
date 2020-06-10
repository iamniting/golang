package main

func getImportance(employees []*Employee, id int) int {
	var getSum func(m map[int]*Employee, id int) int

	getSum = func(m map[int]*Employee, id int) int {
		sum := m[id].Importance
		for _, ID := range m[id].Subordinates {
			sum += getSum(m, ID)
		}
		return sum
	}

	m := map[int]*Employee{}
	for _, e := range employees {
		m[e.Id] = e
	}

	return getSum(m, id)
}
