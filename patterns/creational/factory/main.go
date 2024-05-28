package factory

import "fmt"

// Factory pattern is a creational pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
// Context sample: A factory pattern is used to create different types of digital devices.

func Run() {
	// Create a new smartphone
	smartphone := GetDevice(DeviceType(SmartPhoneType))
	fmt.Println(smartphone.Use())
	laptop := GetDevice(DeviceType(LaptopType))
	fmt.Println(laptop.Use())
	desktop := GetDevice(DeviceType(DesktopType))
	fmt.Println(desktop.Use())
}
