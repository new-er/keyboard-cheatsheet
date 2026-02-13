package nodes

type Node interface {
	Equals(other Node) bool
	ToString() string
}
