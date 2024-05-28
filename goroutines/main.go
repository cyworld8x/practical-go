package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// Go routines are lightweight threads of execution that run concurrently.
// They are used to run multiple tasks simultaneously.
// Go routines are created using the go keyword followed by a function invocation.
// The function runs concurrently in a new go routine.

// Run is a function that demonstrates the use of go routines.
// It creates a go routine that prints a message.
// The main function also prints a message.
// The messages from the go routine and the main function are printed concurrently.
// The order of the messages may vary each time the program is run.
// The go routine may not print the message if the main function completes before the go routine starts.
// The go routine may not print the message if the program exits before the go routine starts.

// waitGroup is a synchronization primitive that is used to wait for a collection of go routines to finish.
// The Add method adds the number of go routines to wait for.
// The Done method decrements the count of go routines that are waiting.
// The waitGroup is used to wait for the go routine to finish before the program exits.

func printMessage(str string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s %d second(s)\n", str, i)
	}

}

func Run() {
	var wg sync.WaitGroup
	printMessage("Direct call: ")

	// Add 1 to the wait group
	// The go routine is started in a new go routine.
	// The wait group is decremented when the go routine finishes.
	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage("Indirect call:")
	}()
	go func(str string) {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}("Hello from a go routine!")

	time.Sleep(1 * time.Second)
	fmt.Println("Hello from main function!")
	// Wait for the go routine to finish
	wg.Wait()

}
