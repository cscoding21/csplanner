package utils

// ValToRef converts a value into a reference to a string
func ValToRef[T any](input T) *T {
	return &input
}

// RefToVal converts a string into a reference to a string
func RefToVal[T any](input *T) T {
	return *input
}

// ValToRefSlice converts a slice into a slice of reference types
func ValToRefSlice[T any](input []T) []*T {
	out := []*T{}

	if len(input) == 0 {
		return out
	}

	for i := range input {
		out = append(out, &input[i])
	}

	return out
}

// RefToValSlice converts a slice of reference types into a slice of values
func RefToValSlice[T any](input []*T) []T {
	out := []T{}

	if len(input) == 0 {
		return out
	}

	for i := range input {
		out = append(out, *input[i])
	}

	return out
}
