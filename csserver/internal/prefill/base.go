package prefill

type Prefill[T any] struct {
	Name        string
	Description string
	Object      T
}
