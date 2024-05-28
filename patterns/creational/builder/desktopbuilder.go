package builder

import "fmt"

// Computer
type DesktopBuilder struct {
	monitor  Monitor
	keyboard KeyBoard
	mouse    Mouse
	graphic  Graphic
}

func (d *DesktopBuilder) New() *DesktopBuilder {
	return &DesktopBuilder{}
}

func (d *DesktopBuilder) setGraphic(graphic Graphic) {
	d.graphic = graphic
}

func (d *DesktopBuilder) setMonitor(monitor Monitor) {
	d.monitor = monitor
}

func (d *DesktopBuilder) setKeyBoard(keyboard KeyBoard) {
	d.keyboard = keyboard
}

func (d *DesktopBuilder) setMouse(mouse Mouse) {
	d.mouse = mouse
}

func (desktopBuilder *DesktopBuilder) Show() {
	fmt.Printf("Desktop with monitor type: %s\n", desktopBuilder.monitor)
	fmt.Printf("Desktop with keyboard type: %s\n", desktopBuilder.keyboard)
	fmt.Printf("Desktop with mouse type: %s\n", desktopBuilder.mouse)
	fmt.Printf("Desktop with graphic type: %s\n", desktopBuilder.graphic)
}

func (desktopBuilder *DesktopBuilder) Build() *Computer {
	desktopBuilder.setMonitor(Monitor(LCD))
	desktopBuilder.setKeyBoard(KeyBoard(USB))
	desktopBuilder.setMouse(Mouse(USB))
	desktopBuilder.setGraphic(Graphic(IntegratedCard))
	return &Computer{
		monitor:  desktopBuilder.monitor,
		keyboard: desktopBuilder.keyboard,
		mouse:    desktopBuilder.mouse,
		graphic:  desktopBuilder.graphic,
		ShowSpec: desktopBuilder.Show,
	}
}
