package prototype

import (
	"fmt"
)

// Leaf struct
type Leaf struct {
	name string
}

func (l *Leaf) ls(indentation string) {
	fmt.Println(indentation + l.name)
}

func (l *Leaf) clone() INode {
	return &Leaf{name: l.name + "_clone"}
}
