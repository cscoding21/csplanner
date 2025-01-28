package common

// IsOneOf test to see if the input is equal to a list of items
func IsOneOf[T comparable](input T, opts ...T) bool {
	for _, o := range opts {
		if input == o {
			return true
		}
	}

	return false
}
