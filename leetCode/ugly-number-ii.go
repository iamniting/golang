package main

import "math"

func nthUglyNumber(n int) int {

	ugly := make([]int, n)
	ugly[0] = 1

	var i2, i3, i5 int
	nextMultipleOf2, nextMultipleOf3, nextMultipleOf5 := 2, 3, 5

	for i := 1; i < n; i++ {
		ugly[i] = int(
			math.Min(float64(nextMultipleOf2),
				math.Min(float64(nextMultipleOf3), float64(nextMultipleOf5)),
			),
		)

		if ugly[i] == nextMultipleOf2 {
			i2++
			nextMultipleOf2 = ugly[i2] * 2
		}

		if ugly[i] == nextMultipleOf3 {
			i3++
			nextMultipleOf3 = ugly[i3] * 3
		}

		if ugly[i] == nextMultipleOf5 {
			i5++
			nextMultipleOf5 = ugly[i5] * 5
		}
	}

	return ugly[n-1]
}

/*`
Example:
Let us see how it works

initialize
    ugly[] = | 1 |
    i2 = i3 = i5 = 0;

First iteration
    ugly[1] = Min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
	    = Min(2, 3, 5)
	    = 2
    ugly[] = | 1 | 2 |
    i2 = 1, i3 = i5 = 0 ( i2 got incremented )

Second iteration
    ugly[2] = Min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
            = Min(4, 3, 5)
            = 3
    ugly[] =  | 1 | 2 | 3 |
    i2 = 1, i3 = 1, i5 = 0 ( i3 got incremented )

Third iteration
    ugly[3] = Min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
            = Min(4, 6, 5)
            = 4
    ugly[] = | 1 | 2 | 3 |  4 |
    i2 = 2, i3 = 1, i5 = 0 ( i2 got incremented )

Fourth iteration
    ugly[4] = Min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
            = Min(6, 6, 5)
            = 5
    ugly[] = | 1 | 2 | 3 |  4 | 5 |
    i2 = 2, i3 = 1, i5 = 1 ( i5 got incremented )

Fifth iteration
    ugly[4] = Min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
            = Min(6, 6, 10)
            = 6
    ugly[] = | 1 | 2 | 3 |  4 | 5 | 6 |
    i2 = 3, i3 = 2, i5 = 1 ( i2 and i3 got incremented )

Will continue same way
`*/
