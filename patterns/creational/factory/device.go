package factory

type Device struct {
	Brand string
	Name  string
}

// DigitalDevice interface
type IDevice interface {
	connectInternet() string
	playGame() string
	listenMusic() string
	work() string
	Use() string
}

func (d *Device) connectInternet() string {
	return d.Name + " connects to the internet"
}

func (d *Device) playGame() string {
	return "Play game on " + d.Brand + " " + d.Name
}

func (d *Device) listenMusic() string {
	return "Listen to music on " + d.Name
}

func (d *Device) work() string {
	return "Working by using " + d.Name
}

func (d *Device) Use() string {
	return d.connectInternet() + ", " + d.playGame() + ", " + d.listenMusic() + ", " + d.work()
}

type DeviceType string

const (
	SmartPhoneType DeviceType = "smartphone"
	LaptopType     DeviceType = "laptop"
	DesktopType    DeviceType = "desktop"
)
