package bridge

import "fmt"

// Window struct
type Window struct {
	printer Printer
}

// Print method
func (w *Window) Print() string {
	fmt.Println("Print request for windows")
	return w.printer.PrintFile()
}

// SetPrinter method
func (w *Window) SetPrinter(p Printer) {
	w.printer = p
}
