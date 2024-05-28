package abstractfactory

// DellKeyboard struct
type DellKeyboard struct {
	Keyboard
}

// DellMonitor struct
type DellMonitor struct {
	Monitor
}

// Dell struct
type Dell struct{}

func (d *Dell) createKeyboard() IKeyboard {
	return &DellKeyboard{
		Keyboard: Keyboard{
			keys: "hello Dell keyboard",
		},
	}
}

func (d *Dell) createMonitor() IMonitor {
	return &DellMonitor{
		Monitor: Monitor{
			screenSize: "1920x1080",
		},
	}
}
