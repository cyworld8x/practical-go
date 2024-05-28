package factory

// Desktop struct
type Desktop struct {
	Device
}

func newDesktop() IDevice {
	return &Desktop{
		Device: Device{
			Brand: "Apple",
			Name:  "iMac",
		},
	}
}
