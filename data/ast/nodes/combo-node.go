package nodes

import (
	"fmt"
)

type ComboNode struct {
	Left  Node
	Right Node
}

func NewComboNode(left Node, right Node) ComboNode {
	return ComboNode{Left: left, Right: right}
}

func (c ComboNode) Equals(other Node) bool {
	if otherCombo, ok := other.(ComboNode); ok {
		return c.Left.Equals(otherCombo.Left) && c.Right.Equals(otherCombo.Right)
	}
	return false
}
func (c ComboNode) ToString() string {
	return fmt.Sprintf("%s+%s", c.Left.ToString(), c.Right.ToString())
}
