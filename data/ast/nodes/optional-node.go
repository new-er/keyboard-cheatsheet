package nodes

type OptionalNode struct {
	Key Node
}

func NewOptionalNode(key Node) OptionalNode {
	return OptionalNode{Key: key}
}

func (o OptionalNode) Equals(other Node) bool {
	if otherOptional, ok := other.(OptionalNode); ok {
		return o.Key.Equals(otherOptional.Key)
	}
	return false
}
func (o OptionalNode) ToString() string {
	return o.Key.ToString() + "?"
}
