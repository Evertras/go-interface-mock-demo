package main

import (
	"github.com/Evertras/go-interface-mock-demo/pkg/external/concrete"
	"github.com/Evertras/go-interface-mock-demo/internal/server"
)

func main() {
	// Just to be very explicit for this demo, we're declaring this up front.
	// This is a concrete instance of a struct, not an interface.
	var concreteClient *concrete.ConcreteClient
	concreteClient = concrete.New()

	// Here's the magic.  Even though the server package has no idea what the
	// concrete package is, and the concrete package has no idea what the server
	// package is, the *ConcreteClient type satisfies the interface so it's accepted.
	server := server.New(concreteClient)

	// Normally we'd start listening and all that, but just to demonstrate
	// that it's working we'll call a method directly that uses the underlying
	// client.
	server.HackStuff()
}
