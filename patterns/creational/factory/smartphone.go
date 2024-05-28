package factory

// Smartphone struct
type SmartPhone struct {
	Device
}

func newSmartPhone() IDevice {
	return &SmartPhone{
		Device: Device{
			Brand: "Apple",
			Name:  "iPhone 15",
		},
	}
}
