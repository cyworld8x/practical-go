package abstractfactory

// IMonitor interface
type IMonitor interface {
	display() string
	resolution(size string)
}

// Monitor struct
type Monitor struct {
	screenSize string
}

func (m *Monitor) display() string {
	return "Displaying on resolution " + m.screenSize
}

func (m *Monitor) resolution(size string) {
	m.screenSize = size
}
