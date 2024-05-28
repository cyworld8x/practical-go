package generics

//Starting with version 1.18, Go has added support for generics, also known as type parameters.
//Generics allow you to write functions, data structures, and algorithms that can work with any type.

//A type parameter is a placeholder for a type that is specified when the function or data structure is used.
//Type parameters are enclosed in square brackets [] after the function name or data structure name.
//Type parameters are declared using the keyword type followed by the type constraint.
//Type constraints specify the type of values that can be used with the type parameter.
//Type constraints can be any, interface{}, or a specific type.
//The any type constraint allows any type to be used with the type parameter.

import (
	"fmt"
)

type List[T any] struct {
	head, tail *item[T]
}

type item[T any] struct {
	next  *item[T]
	value T
}

func (d *List[T]) Add(value T) {
	newItem := &item[T]{value: value}
	if d.head == nil {
		d.head = newItem
	} else {
		d.tail.next = newItem
	}
	d.tail = newItem
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (d List[T]) Keys() []T {
	keys := []T{}
	for i := d.head; i != nil; i = i.next {
		keys = append(keys, i.value)
	}
	return keys
}

func Run() {
	listInts := List[int]{}
	listInts.Add(1)
	listInts.Add(2)
	fmt.Println(listInts.Keys())

	listStrings := List[string]{}
	listStrings.Add("a")
	listStrings.Add("b")

	fmt.Println(listStrings.Keys())

	var m = map[int]string{1: "2", 3: "4", 5: "8", 9: "10", 11: "12"}
	fmt.Println("keys:", MapKeys(m))

}
