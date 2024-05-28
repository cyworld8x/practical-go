package composite

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
