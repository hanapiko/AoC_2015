package floor

func FloorNum(s string) int{
	count := 0
	for _, ch := range s{
		if ch == '(' {
			count++
			// continue
		} else {
		count--
	}
	}
	return count
}

