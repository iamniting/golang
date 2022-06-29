package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}
	for i := 50; i > 0; i = i - 1 {
		s = append(s, i)
	}

	fmt.Println("Before sorting")
	fmt.Println(s)
	mergeSort(s, 0, len(s)-1)
	fmt.Println("After sorting")
	fmt.Println(s)
}

func mergeSort(s []int, left, right int) {
	if left < right {

		var wg sync.WaitGroup
		wg.Add(2)

		mid := (left + right) / 2

		go func() {
			defer wg.Done()
			mergeSort(s, left, mid)
		}()

		go func() {
			defer wg.Done()
			mergeSort(s, mid+1, right)
		}()

		wg.Wait()

		merge(s, left, right)
	}
}

func merge(s []int, leftStart, rightEnd int) {

	temp := []int{}

	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	left, right := leftStart, rightStart

	// copy elements in sorted fashion into temp array as long as both arrays have something
	for left <= leftEnd && right <= rightEnd {
		if s[left] <= s[right] {
			temp = append(temp, s[left])
			left++
		} else if s[left] > s[right] {
			temp = append(temp, s[right])
			right++
		}
	}

	// copy elements into temp array as long as left array is not done
	for left <= leftEnd {
		temp = append(temp, s[left])
		left++
	}

	// copy elements into temp array as long as right array is not done
	for right <= rightEnd {
		temp = append(temp, s[right])
		right++
	}

	// copy temp array elements into the main array
	for _, n := range temp {
		s[leftStart] = n
		leftStart++
	}
}
