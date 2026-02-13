package nodes

type AliasNode struct {
	Alias string
	Key   Node
}

func NewAliasNode(key Node, alias string) AliasNode {
	return AliasNode{Alias: alias, Key: key}
}

func (a AliasNode) Equals(other Node) bool {
	if otherAlias, ok := other.(AliasNode); ok {
		return a.Alias == otherAlias.Alias && a.Key.Equals(otherAlias.Key)
	}
	return false
}
func (a AliasNode) ToString() string {
	return a.Alias
}
