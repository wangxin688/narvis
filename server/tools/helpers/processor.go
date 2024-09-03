package helpers

func EnsureSliceNotNil[T any](ptr *[]T) []T {
	if ptr == nil {
		var emptySlice []T
		return emptySlice
	}
	return *ptr
}
