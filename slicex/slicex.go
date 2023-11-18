package slicex

func Contains[T comparable](s []T, t T) bool {
	for i := range s {
		if s[i] == t {
			return true
		}
	}
	return false
}
