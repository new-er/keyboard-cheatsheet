package nodes

type KeyNode struct {
	Name string
}

func NewKeyNode(name string) KeyNode {
	return KeyNode{Name: name}
}

func (k KeyNode) Equals(other Node) bool {
	if otherKey, ok := other.(KeyNode); ok {
		return k.Name == otherKey.Name
	}
	return false
}
func (k KeyNode) ToString() string {
	return k.Name
}
