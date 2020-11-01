package main

func numTeams(rating []int) int {

	var res int

	for j, n := range rating {
		var lsmall, llarge, rsmall, rlarge int

		for i := 0; i < j; i++ {
			if rating[i] < n {
				lsmall++
			} else if rating[i] > n {
				llarge++
			}
		}

		for k := j + 1; k < len(rating); k++ {
			if rating[k] < n {
				rsmall++
			} else if rating[k] > n {
				rlarge++
			}
		}

		// (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k])
		res += (lsmall * rlarge) + (llarge * rsmall)
	}

	return res
}
