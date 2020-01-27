package concrete

import "fmt"

// ConcreteClient does some interesting thing that is useful, but is a hard
// dependency.  We can't mock it with this package alone because there is no
// interface provided for us.
type ConcreteClient struct{}

// HackServers hits some remote servers and hacks them, just like the movies!
func (c *ConcreteClient) HackServers() error {
	fmt.Println("I'm hacking your servers!")

	// Obviously nothing could ever go wrong
	return nil
}

// MineBitcoins will mine some bitcoins for us
func (c *ConcreteClient) MineBitcoins() {
	fmt.Println("I'm mining bitcoins!")
}

// New returns a concrete client ready to hack servers and mine bitcoins.  Again,
// note that this returns a concrete struct, not an interface.
func New() *ConcreteClient {
	return &ConcreteClient{}
}
