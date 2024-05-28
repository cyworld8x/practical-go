package singleton

import (
	"fmt"
	"sync"
)

//singleton pattern is a creational design pattern that ensures that a class has only one instance and provides a global point of access to that instance.
//The singleton pattern is a design pattern that restricts the instantiation of a class to one object.
//This is useful when exactly one object is needed to coordinate actions across the system.

type Monitor string
type KeyBoard string
type Mouse string
type Graphic string

type Computer struct {
	monitor  Monitor
	keyboard KeyBoard
	mouse    Mouse
	graphic  Graphic
}

// sync.Mutex is a mutual exclusion lock (mutex) data structure that is used to protect shared data from being simultaneously accessed or modified by multiple threads.
var lock = &sync.Mutex{}

var singleComputorInstance *Computer

func GetInstance(i int) *Computer {

	if singleComputorInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleComputorInstance == nil {
			fmt.Printf("Creating Single Computer instance #%d Now\n", i)
			singleComputorInstance = &Computer{
				monitor:  Monitor("LED"),
				keyboard: KeyBoard("Wireless"),
				mouse:    Mouse("Wireless"),
				graphic:  Graphic("DedicatedCard"),
			}
			return singleComputorInstance
		}
	}
	fmt.Printf("Single computer instance #%d already created.\n", i)
	return singleComputorInstance
}

func Run() {
	for i := 0; i < 10; i++ {
		go GetInstance(i)
	}
	fmt.Scanln()
}
