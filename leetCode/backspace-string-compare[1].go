package main

func backspaceCompare(S string, T string) bool {

	i, j := len(S)-1, len(T)-1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {

		for i >= 0 {
			if S[i] == '#' {
				skipS++
			} else if skipS > 0 {
				skipS--
			} else {
				break
			}
			i--
		}

		for j >= 0 {
			if T[j] == '#' {
				skipT++
			} else if skipT > 0 {
				skipT--
			} else {
				break
			}
			j--
		}

		if i >= 0 && j >= 0 && S[i] == T[j] {
			i--
			j--
		} else {
			return i == -1 && j == -1
		}

	}
	return true
}
