package closure

//Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables.
// The function value keeps a reference to the variables from the enclosing function.
// The function value is bound to the variables from the enclosing function.

import "fmt"

func Func() func(s string) {
	return func(str string) {
		fmt.Println(str)
	}
}

func Run() {

	fmt.Println("[Closure] ----------> Start Run print := Func() print(s string)  <----------")
	print := Func()
	print("Hello, World!")
}
