package abstractfactory

import "fmt"

// abtract factory pattern is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes.
// The abstract factory pattern is useful when you have a family of related or dependent objects that need to be created together.
// Context example: An abstract factory pattern is used to create a family of related digital devices, such as a keyboard and a monitor.
// Here are two brands of digital devices: HP and Dell.
// Each brand has its own factory that creates its own keyboard and monitor.
// The getFactory function is used to get the factory of a specific brand.
// The interface factory has two methods: createKeyboard and createMonitor.
// The createKeyboard method creates a keyboard of a specific brand, and the createMonitor method creates a monitor of a specific brand.
// The HP struct is a factory that creates HP keyboards and monitors.
// The HPKeyboard struct is a keyboard of the HP brand.
// The HPMonitor struct is a monitor of the HP brand.
// The Dell struct is a factory that creates Dell keyboards and monitors.
// The DellKeyboard struct is a keyboard of the Dell brand.
// The DellMonitor struct is a monitor of the Dell brand.
// The Run function is used to create a keyboard and a monitor of each brand and display the typing and display information.

func Run() {
	factory := getFactory("HP")
	keyboard := factory.createKeyboard()
	fmt.Printf("%s %T \n", keyboard.typing(), factory)
	monitor := factory.createMonitor()
	fmt.Println(monitor.display())

	factory = getFactory("Dell")
	keyboard = factory.createKeyboard()
	fmt.Printf("%s %T \n", keyboard.typing(), factory)
	monitor = factory.createMonitor()
	fmt.Println(monitor.display())
}
