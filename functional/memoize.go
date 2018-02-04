package functional

func Memoize(fn func(int) int) (func(int) int, map[int]int) {
	cache := make(map[int]int)

	return func(i int) int {
		if cachedValue, isCached := cache[i]; isCached {
			return cachedValue
		}

		result := fn(i)
		cache[i] = result
		return result
	}, cache
}
