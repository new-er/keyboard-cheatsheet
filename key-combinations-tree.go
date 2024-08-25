package main

import (
	"fmt"
	"strings"
)

type Node map[string]interface{}

func (n *Node) Print() {
	n.print(0)
}

func (n *Node) print(depth int) {
	for k, v := range *n {
		switch v := v.(type) {
		case Node:
			fmt.Println(strings.Repeat(" ", depth), "∟", k)
			v.print(depth + 1)
		case string:
			fmt.Println(strings.Repeat(" ", depth), "∟", k, ":", v)
		}
	}
}

func NewKeyCombinationsTree() *Node {
	root := Node{
		"windows": Node{
			"ALT": Node{
				"TAB": "Switch between open apps",
			},
		},
		"firefox": Node{
			"CTRL": Node{
				"t": "Open a new tab",
			},
		},
		"PowerShell": Node{
			"CTRL": Node{
				"SHIFT": Node{
					"TAB": "Create a new tab",
				},
			},
		},
	}
	return &root
}

func (n *Node) filterByPrograms(programs []string) *Node {
	node := Node{}
	for k, v := range *n {
		for _, program := range programs {
			if strings.Contains(program, k) {
				node[k] = v
			}
		}
	}
	return &node

}
