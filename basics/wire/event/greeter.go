package event

import "time"

// NewGreeter initializes a Greeter. If the current epoch time is an even
// number, NewGreeter will create a grumpy Greeter.
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greeter is the type charged with greeting guests.
type Greeter struct {
	Grumpy  bool
	Message Message
}

// Greet produces a greeting for guests.
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}
