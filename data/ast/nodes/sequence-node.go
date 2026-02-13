package nodes

import (
	"fmt"
)

type SequenceNode struct {
	Left  Node
	Right Node
}

func NewSequenceNode(left Node, right Node) SequenceNode {
	return SequenceNode{Left: left, Right: right}
}

func (s SequenceNode) Equals(other Node) bool {
	if otherSequence, ok := other.(SequenceNode); ok {
		return s.Left.Equals(otherSequence.Left) && s.Right.Equals(otherSequence.Right)
	}
	return false
}
func (s SequenceNode) ToString() string {
	return fmt.Sprintf("%s %s", s.Left.ToString(), s.Right.ToString())
}
