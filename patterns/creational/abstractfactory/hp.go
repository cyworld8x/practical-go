package abstractfactory

// HP struct

type HPKeyboard struct {
	Keyboard
}

type HPMonitor struct {
	Monitor
}

type HP struct{}

func (h *HP) createKeyboard() IKeyboard {
	return &HPKeyboard{
		Keyboard: Keyboard{
			keys: "hello HP keyboard",
		},
	}
}

func (h *HP) createMonitor() IMonitor {
	return &HPMonitor{
		Monitor: Monitor{
			screenSize: "1024x768",
		},
	}
}
