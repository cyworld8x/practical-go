package bridge

type Computer interface {
	Print() string
	SetPrinter(p *Printer)
}
