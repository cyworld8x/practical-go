package bridge

import "fmt"

// bridge pattern is about preferring composition over inheritance.
// Implementation details are pushed from a hierarchy to another object with a separate hierarchy.
// The bridge pattern is useful when both the class and what it does vary often.
// The class itself can be thought of as the abstraction and what the class can do as the implementation.
// The bridge pattern can also be thought of as two layers of abstraction.
// In this example, the bridge pattern is demonstrated by having two hierarchies: one for the type of computer and one for the type of printer.
// The computer hierarchy is implemented by the Computer interface, and the printer hierarchy is implemented by the Printer interface.
// The Mac and Window structs implement the Computer interface, and the Epson and HP structs implement the Printer interface.
// The Mac and Window structs have a printer field of type Printer, and the Print method prints a message and calls the PrintFile method on the printer field.
// The SetPrinter method sets the printer field.
// The main.go file demonstrates the bridge pattern by creating a Mac and a Window, creating an Epson and an HP, setting the printer for the Mac and Window, and calling the Print method on the Mac and Window.

func Run() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	mac := &Mac{}
	mac.SetPrinter(hpPrinter)
	fmt.Println(mac.Print())
	mac.SetPrinter(epsonPrinter)
	fmt.Println(mac.Print())

	window := &Window{}
	window.SetPrinter(epsonPrinter)
	fmt.Println(window.Print())
	window.SetPrinter(hpPrinter)
	fmt.Println(window.Print())
}
