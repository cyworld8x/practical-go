package prototype

import "fmt"

type Branch struct {
	nodes []INode
	name  string
}

func (b *Branch) ls(indentation string) {
	fmt.Println(indentation + b.name)
	for _, node := range b.nodes {
		node.ls(indentation + indentation)
	}
}

func (b *Branch) clone() INode {
	cloneBranch := &Branch{name: b.name + "_clone"}
	for _, node := range b.nodes {
		cloneBranch.addNode(node.clone())
	}
	return cloneBranch
}

func (b *Branch) addNode(node INode) {
	b.nodes = append(b.nodes, node)
}
