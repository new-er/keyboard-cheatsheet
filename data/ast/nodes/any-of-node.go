package nodes

import (
	"fmt"
)

type AnyOfNode struct {
	Left  Node
	Right Node
}

func NewAnyOfNode(left Node, right Node) AnyOfNode {
	return AnyOfNode{Left: left, Right: right}
}

func (a AnyOfNode) Equals(other Node) bool {
	if otherAnyOf, ok := other.(AnyOfNode); ok {
		return a.Left.Equals(otherAnyOf.Left) && a.Right.Equals(otherAnyOf.Right)
	}
	return false
}
func (a AnyOfNode) ToString() string {
	return fmt.Sprintf("%s|%s", a.Left.ToString(), a.Right.ToString())
}
