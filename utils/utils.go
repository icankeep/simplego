package utils

func EmptyOrDefault[T any](s []T, defaultValue T) (t T) {
	if len(s) == 0 {
		return
	}
	return defaultValue
}

func SafeIndexValueOrDefault[T any](s []T, idx int) (t T) {
	if len(s) <= idx {
		return
	}
	return s[idx]
}

func If[T any](cond bool, onTrue T, onFalse T) T {
	if cond {
		return onTrue
	}
	return onFalse
}
