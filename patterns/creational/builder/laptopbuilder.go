package builder

import "fmt"

// Computer
type LaptopBuilder struct {
	monitor  Monitor
	keyboard KeyBoard
	mouse    Mouse
	graphic  Graphic
}

func (l *LaptopBuilder) New() *LaptopBuilder {
	return &LaptopBuilder{}
}

func (l *LaptopBuilder) setGraphic(graphic Graphic) {
	l.graphic = graphic
}

func (l *LaptopBuilder) setMonitor(monitor Monitor) {
	l.monitor = monitor
}

func (l *LaptopBuilder) setKeyBoard(keyboard KeyBoard) {
	l.keyboard = keyboard
}

func (l *LaptopBuilder) setMouse(mouse Mouse) {
	l.mouse = mouse
}

func (laptopBuilder *LaptopBuilder) Show() {
	fmt.Printf("Laptop with monitor type: %s\n", laptopBuilder.monitor)
	fmt.Printf("Laptop with keyboard type: %s\n", laptopBuilder.keyboard)
	fmt.Printf("Laptop with mouse type: %s\n", laptopBuilder.mouse)
	fmt.Printf("Laptop with graphic type: %s\n", laptopBuilder.graphic)
}

func (laptopBuilder *LaptopBuilder) Build() *Computer {

	laptopBuilder.setMonitor(Monitor(LED))
	laptopBuilder.setKeyBoard(KeyBoard(Wireless))
	laptopBuilder.setMouse(Mouse(Wireless))
	laptopBuilder.setGraphic(Graphic(DedicatedCard))
	return &Computer{
		monitor:  laptopBuilder.monitor,
		keyboard: laptopBuilder.keyboard,
		mouse:    laptopBuilder.mouse,
		graphic:  laptopBuilder.graphic,
		ShowSpec: laptopBuilder.Show,
	}
}
