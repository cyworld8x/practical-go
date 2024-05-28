package builder

//Builder pattern is a creational design pattern that lets you construct complex objects step by step.
//The pattern allows you to produce different types and representations of an object using the same construction code.
//The builder pattern is a design pattern that allows for the step-by-step creation of complex objects using the correct sequence of actions.
//The construction is controlled by a director object that only needs to know the type of object it is to create.
//The director is not responsible for the creation of the object itself, but only for the order in which the builder constructs the object.
//The builder pattern is a design pattern that allows for the step-by-step creation of complex objects using the correct sequence of actions.

// The builder pattern includes 4 main components:
// 1. Product: The final object that is constructed by the builder. It is the object that the builder constructs. Ex: Laptop, Desktop, etc.
// 2. Builder: An interface that specifies how to build parts of the product. Ex: IComputerBuilder.
// 3. ConcreteBuilder: A concrete implementation of the Builder interface that constructs the product. Ex: LaptopBuilder, DesktopBuilder, etc
// 4. Director: The object that controls the construction process by using the builder. Ex: Director

func Run() {
	// Create a laptop
	laptopBuilder := GetComputerBuilder(Laptop)
	desktopBuilder := GetComputerBuilder(Desktop)
	director := NewDirector(laptopBuilder)
	laptop := director.buildComputer()
	laptop.ShowSpec()
	director.setBuilder(desktopBuilder)
	desktop := director.buildComputer()
	desktop.ShowSpec()

}
