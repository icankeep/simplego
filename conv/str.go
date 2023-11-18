package conv

import "strconv"

func Int64Default(s string, defaultValue int64) int64 {
	if len(s) == 0 {
		return defaultValue
	}
	number, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return number
}
