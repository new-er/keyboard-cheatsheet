package helpers

func Where[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Single[T any](ss []T, test func(T) bool) (ret T, found bool) {
	for _, s := range ss {
		if test(s) {
			if found {
				return
			}
			ret = s
			found = true
		}
	}
	return
}

func Map[T, U any](ss []T, f func(T) U) (ret []U) {
	for _, s := range ss {
		ret = append(ret, f(s))
	}
	return
}

func StringSliceEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TakeFirst[T any](ss []T, n int) []T {
	if n > len(ss) {
		return ss
	}
	return ss[:n]
}

func TakeLast[T any](ss []T, n int) []T {
	if n > len(ss) {
		return ss
	}
	return ss[len(ss)-n:]
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if !test(s) {
			ret = append(ret, s)
		}
	}
	return
}
