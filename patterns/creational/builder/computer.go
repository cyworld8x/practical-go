package builder

type Graphic string

const (
	DedicatedCard  Graphic = "DedicatedCard"
	IntegratedCard Graphic = "IntegratedCard"
)

type Monitor string

const (
	LCD Monitor = "LCD"
	LED Monitor = "LED"
)

type KeyBoard string

const (
	QWERTY    KeyBoard = "QWERTY"
	Wired     KeyBoard = "Wired"
	Ergonomic KeyBoard = "Ergonomic"
	Wireless  KeyBoard = "Wireless"
	USB       KeyBoard = "USB"
	Bluetooth KeyBoard = "Bluetooth "
)

type Mouse string

const (
	WiredMouse     Mouse = "WiredMouse"
	WirelessMouse  Mouse = "WirelessMouse"
	BluetoothMouse Mouse = "BluetoothMouse"
)

// Computer
type Computer struct {
	monitor  Monitor
	keyboard KeyBoard
	mouse    Mouse
	graphic  Graphic
	ShowSpec func()
}

type Device int

const (
	Laptop  Device = 1
	Desktop Device = 2
)
