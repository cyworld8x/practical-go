package abstractfactory

// IComputerFactory interface
type IComputerFactory interface {
	createKeyboard() IKeyboard
	createMonitor() IMonitor
}

func getFactory(brand string) IComputerFactory {
	switch brand {
	case "HP":
		return &HP{}
	case "Dell":
		return &Dell{}
	default:
		return nil
	}
}
