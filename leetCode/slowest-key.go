package main

func slowestKey(releaseTimes []int, keysPressed string) byte {

	var resKey byte
	var resTime, prevReleaseTime int

	for i := range keysPressed {

		time := releaseTimes[i] - prevReleaseTime

		if time > resTime || (time == resTime && resKey < keysPressed[i]) {
			resTime = time
			resKey = keysPressed[i]
		}

		prevReleaseTime = releaseTimes[i]
	}

	return resKey
}
