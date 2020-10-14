package main

import "fmt"

func main() {

	/*message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()*/

	event := InitEvent()
	event.Start()


}

type Message string

type Greeter struct {
	Message
}

type Event struct {
	Greeter
}

func NewMessage() Message {
	return Message("this is message")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
