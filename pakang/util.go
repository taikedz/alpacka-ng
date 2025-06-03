package pakang

func ArrayHas(term string, stuff []string) bool {
	for _, thing := range(stuff) {
		if term == thing {
			return true
		}
	}
	return false
}