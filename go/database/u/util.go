package u

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func Ptr[T any](val T) *T {
	return &val
}
