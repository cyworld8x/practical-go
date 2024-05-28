package channels

// Go channels are used to communicate between go routines.
// A channel is a typed conduit through which you can send and receive values with the channel operator, <-.
// The zero value of a channel is nil.
// A nil channel is not usable until it is initialized.
// A channel can be initialized using the make function.
// The make function creates a channel and returns a reference to it.

// unbuffered channels are synchronous.
// A send operation on an unbuffered channel blocks until a receive operation is ready.
// A receive operation on an unbuffered channel blocks until a send operation is ready.

// buffered channels are asynchronous.
// A send operation on a buffered channel blocks only when the buffer is full.
// A receive operation on a buffered channel blocks only when the buffer is empty.
// the order of messages sent and received on a channel is preserved.
// the order of messages sent and received on a channel is not guaranteed.
// the sequence of messages sent and received on a channel is not sequential.

// A channel can be closed using the close function.

// Run is a function that demonstrates the use of channels.
// It creates a channel and sends a message to the channel.
// It receives the message from the channel and prints it.
// The channel is closed after sending the message.
// The receive operation on a closed channel returns the zero value of the channel type.
// The receive operation on a closed channel does not block.

// select statement is used to wait on multiple channel operations.
// The select statement blocks until one of its cases can run.
// The select statement chooses one at random if multiple cases are ready.

// Note: Basic sends and receives on channels are blocking.
// However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.

import (
	"fmt"
	"time"
)

func Run() {
	// Create a unbufferedchannel of type string
	unbufferedchannel := make(chan string)

	// Send a message to the channel
	go func() {
		time.Sleep(1 * time.Second)
		unbufferedchannel <- "Hello channel unbuffered channel, from a go #1 routine!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		unbufferedchannel <- "Hello channel unbuffered channel, from a go #2 routine!"
	}()

	// Receive the message from the channel
	msg := <-unbufferedchannel
	fmt.Println(msg)

	//Close the channel after sending the message
	//close(channel)

	// Receive the message from the channel
	// throws panic: send on closed channel
	//msg = <-channel
	msg = <-unbufferedchannel
	fmt.Println(msg)

	// Create a buffered channel of type string, size 4
	// if the size is lower than the number of messages sent, it will throw a deadlock
	bufferedchannel := make(chan string, 5)

	go func() {
		bufferedchannel <- "Hello buffered channel, from a go routine, send #1 message!"
	}()

	go func() {
		bufferedchannel <- "Hello buffered channel, from a go routine, send #2 message!"
	}()

	// Receive the message from the channel
	bufferedchannel <- "Hello buffered channel, send #3 message!"
	bufferedchannel <- "Hello buffered channel, send #4 message!"
	msg = <-bufferedchannel
	fmt.Println(msg)
	msg = <-bufferedchannel
	fmt.Println(msg)
	msg = <-bufferedchannel
	fmt.Println(msg)
	msg = <-bufferedchannel
	fmt.Println(msg)

	// Build a buffered channel in synchronous mode
	// The order of messages sent and received on a channel is preserved.
	bufferedchannel = make(chan string, 1)

	go func(bufferedchannel chan string) {
		fmt.Print("Hello buffered channel in synchronous mode, from a go routine!")
		time.Sleep(1 * time.Second)
		fmt.Print("send #1 message! \n")
		bufferedchannel <- "#1 message"
	}(bufferedchannel)
	// If remove <-bufferedchannel, the program would exit before the go routines finish
	<-bufferedchannel

	// Build channel direction
	// First channel can be declared as publiser, and second channel as subscriber

	// Create a channel with a direction

	publisher := make(chan string, 1)
	subscriber := make(chan string, 1)

	publish := func(pub chan<- string, msg string) {
		pub <- msg
	}

	subscribe := func(pub <-chan string, sub chan<- string) {
		msg := <-pub
		sub <- msg
	}

	publish(publisher, "Hello from publisher!")
	subscribe(publisher, subscriber)

	fmt.Println(<-subscriber)

	// Build a select statement
	// The select statement is used to wait on multiple channel operations.

	// Create two channels

	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channelOne <- "Hello from channel one! Send message #1 (1 second) at " + time.Now().String()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "Hello from channel one! Send message #2 (3 seconds) at" + time.Now().String()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "Hello from channel one! Send message #3 (2 seconds) at" + time.Now().String()
	}()

	//Now we want to receive a message from either channel one or channel two
	//We can use the select statement to wait on multiple channel operations after 2 seconds

	fmt.Println(time.Now())
	for i := 0; i < 3; i++ {
		select {
		case msg := <-channelOne:
			fmt.Println(msg)
		case msg := <-channelTwo:
			fmt.Println(msg)
		}
	}

	select {
	case msg := <-channelOne:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! No message received in 3 seconds")
	}
	fmt.Println(time.Now())

}
