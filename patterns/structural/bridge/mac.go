package bridge

import "fmt"

// Mac struct
type Mac struct {
	printer Printer
}

// Print method
func (m *Mac) Print() string {
	fmt.Println("Print request for Mac")
	return m.printer.PrintFile()
}

// SetPrinter method
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
