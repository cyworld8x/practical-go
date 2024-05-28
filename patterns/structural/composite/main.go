package composite

// composite pattern is a structural design pattern that lets you compose objects into tree structures to represent part-whole hierarchies.
// Composite lets clients treat individual objects and compositions of objects uniformly.

// Sample usage of the composite pattern
// root
//   branch1
//     branch2
//         leaf 5
//     leaf 3
//     leaf 4
//   leaf 1
//   leaf 2

func Run() {

	//Create a root branch
	root := &Branch{name: "root"}
	//Create a leaf
	leaf1 := &Leaf{name: "leaf 1"}
	//Create a leaf
	leaf2 := &Leaf{name: "leaf 2"}
	//Create a branch
	branch1 := &Branch{name: "branch1"}
	//Create a leaf
	leaf3 := &Leaf{name: "leaf 3"}
	leaf4 := &Leaf{name: "leaf 4"}
	//create a branch 2
	branch2 := &Branch{name: "branch2"}
	//Create a leaf
	leaf5 := &Leaf{name: "leaf 5"}
	branch2.addNode(leaf5)
	branch1.addNode(branch2)
	//Add leafs to branch1
	branch1.addNode(leaf3)
	branch1.addNode(leaf4)

	root.addNode(branch1)
	root.addNode(leaf1)
	root.addNode(leaf2)
	root.ls(" ")
}
