package conversion

import "strconv"

func ToString(integer int) string {
	return strconv.FormatInt(int64(integer), 10)
}

func ToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
