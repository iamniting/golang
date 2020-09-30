package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	var count int

	for i := range flowerbed {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}

	return count >= n
}
