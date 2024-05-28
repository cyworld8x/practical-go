package builder

// ComputerBuilder
type IComputerBuilder interface {
	setMonitor(monitor Monitor)
	setKeyBoard(keyboard KeyBoard)
	setMouse(mouse Mouse)
	setGraphic(graphic Graphic)
	Build() *Computer
}

func GetComputerBuilder(device Device) IComputerBuilder {

	if device == Laptop {
		return &LaptopBuilder{}
	} else if device == Desktop {
		return &DesktopBuilder{}
	}
	return nil
}
