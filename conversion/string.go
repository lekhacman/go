package conversion

import "strconv"

func ToString(integer int) string {
	return strconv.FormatInt(int64(integer), 10)
}
