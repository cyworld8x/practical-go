package bridge

// Epson struct
type Epson struct {
}

// PrintFile method
func (d *Epson) PrintFile() string {
	return "Printing by a Epson Printer"
}
