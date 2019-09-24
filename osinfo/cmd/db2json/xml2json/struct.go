package xml2json

import (
	"strings"
)

type Node struct {
	Children              map[string]Nodes
	Data                  string
	ChildrenAlwaysAsArray bool
}

type Nodes []*Node

func (n *Node) AddChild(s string, c *Node) {
	if n.Children == nil {
		n.Children = map[string]Nodes{}
	}
	n.Children[s] = append(n.Children[s], c)
}

func (n *Node) IsComplex() bool { return len(n.Children) > 0 }
func (n *Node) GetChild(path string) *Node {
	result := n
	names := strings.Split(path, ".")
	for _, name := range names {
		children, exists := result.Children[name]
		if !exists {
			return nil
		}
		if len(children) == 0 {
			return nil
		}
		result = children[0]
	}
	return result
}
