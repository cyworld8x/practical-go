package pointer

import (
	"crypto/sha256"
	"fmt"
)

// Go supports pointers, allowing you to pass references to values and records within your program.
// A pointer is a variable that stores the memory address of another variable.
// A pointer is declared using the * symbol followed by the type of the variable it points to.
// The zero value of a pointer is nil.
// A nil pointer does not point to any memory location.
// A pointer can be dereferenced using the * symbol to access the value it points to.
// A pointer can be used to modify the value it points to.
// A pointer can be passed to a function to allow the function to modify the value it points to.
// A pointer can be used to avoid copying large data structures when passing them to functions.
// A pointer can be used to create data structures such as linked lists and trees.
// A pointer can be used to create recursive, cyclic, mutable, shared data structures.

type Vertex struct {
	X int
	Y int
}

func (v Vertex) length() int {
	return v.X*v.X + v.Y*v.Y
}

func (v Vertex) hash() []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", v)))
	bs := h.Sum(nil)
	return bs
}

type Content string

func (c Content) length() int {
	return len(c)
}

func (c Content) hash() []byte {
	h := sha256.New()
	h.Write([]byte(c))
	bs := h.Sum(nil)
	return bs
}

type IMeasure interface {
	length() int
	hash() []byte
}

func measure(m IMeasure) {
	fmt.Printf("Execute m.length() func via interface of %T => length of %T: %d \n", m, m, m.length())
	fmt.Printf("Execute m.hash() func via interface of %T => hash value: %x\n", m, m.hash())
}

func assignValue(value int) {
	value = 100
}

func assignPointerValue(value *int) {
	*value = 100
}

func Run() {

	fmt.Println("[Pointer] ----------> Start Run Pointer value := 10<----------")
	value := 10
	assignValue(value)
	fmt.Println("Print value assignValue(value) => value:", value)

	assignPointerValue(&value)
	fmt.Println("Print pointer value assignValue(&value) => value:", value)

	assignPointerValue(&value)
	fmt.Println("Print pointer address assignValue(value) => &value:", &value)

	fmt.Println("[Pointer] ----------> Start Run Pointer with Vertx(1,2)<----------")
	v := Vertex{1, 2}
	fmt.Println("Print Vertex v := Vertex{1, 2} => v:", v)
	v.X = 4
	fmt.Println("Set Vertex v.X = 4 => v:", v)
	fmt.Println("Set Vertex v.X = 4 => &v:", &v)
	fmt.Println("Get Length of Vertex v.length() => v:", v.length())

	vertex := Vertex{1, 2}
	content := Content("Hello, World!")
	measure(vertex)
	measure(content)

}
