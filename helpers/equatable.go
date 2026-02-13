package helpers

type Equatable interface {
	Equals(other Equatable) bool
}
