package prototype

// INode interface

type INode interface {
	ls(indentation string)
	clone() INode
}
