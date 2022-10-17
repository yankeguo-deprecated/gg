package gg

func Map[T any, U any](s []T, fn func(T) U) []U {
	var o []U
	if s == nil {
		return o
	}
	o = make([]U, len(s), len(s))
	for i := range s {
		o[i] = fn(s[i])
	}
	return o
}

func Set[T comparable](s []T) map[T]struct{} {
	o := make(map[T]struct{})
	for _, k := range s {
		o[k] = struct{}{}
	}
	return o
}

func Keys[T comparable, U any](m map[T]U) []T {
	var o []T
	for k := range m {
		o = append(o, k)
	}
	return o
}

func Values[T comparable, U any](m map[T]U) []U {
	var o []U
	for _, v := range m {
		o = append(o, v)
	}
	return o
}

func Filter[T any](s []T, fn func(T) bool) []T {
	var o []T
	for _, v := range s {
		if fn(v) {
			o = append(o, v)
		}
	}
	return o
}
