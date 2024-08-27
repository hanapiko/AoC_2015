package floor

func FloorNum(s string) int {
	count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
			// continue
		} else {
			count--
		}
	}
	return count
}

func Position(s string) int {
	count := 0
	pos := 0
	for i, ch := range s {
		if ch == '(' {
			count++
		} else {
			count--
		}
		if count == -1 {
			pos = i + 1
			break
		}
	}
	return pos
}
