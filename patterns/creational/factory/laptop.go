package factory

// Laptop struct
type Laptop struct {
	Device
}

func newLaptop() IDevice {
	return &Laptop{
		Device: Device{
			Brand: "Apple",
			Name:  "MacBook Pro",
		},
	}
}
