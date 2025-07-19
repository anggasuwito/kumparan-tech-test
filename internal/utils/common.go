package utils

func InArray(v string, compare []string) bool {
	for _, c := range compare {
		if v == c {
			return true
		}
	}
	return false
}
