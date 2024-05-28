package factory

func useDevice(d IDevice) string {
	return d.connectInternet() + ", " + d.playGame() + ", " + d.listenMusic() + ", " + d.work()
}

func GetDevice(deviceType DeviceType) IDevice {
	switch deviceType {
	case SmartPhoneType:
		return newSmartPhone()
	case LaptopType:
		return newLaptop()
	case DesktopType:
		return newDesktop()
	default:
		return newDesktop()
	}
}
