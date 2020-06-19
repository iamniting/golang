package main

import "math"

func distanceBetweenBusStops(distance []int, start int, destination int) int {

	if start > destination {
		start, destination = destination, start
	}

	sumClockWise, sumAntiClockWise := 0, 0

	for _, d := range distance[start:destination] {
		sumClockWise += d
	}

	for _, d := range append(distance[:start], distance[destination:]...) {
		sumAntiClockWise += d
	}

	return int(math.Min(float64(sumClockWise), float64(sumAntiClockWise)))
}
