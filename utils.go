package utils

func Contains(a string, x []string) bool {
	for _, s := range x {
		if s == a {
			return true
		}
	}
	return false
}
