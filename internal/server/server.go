package server

import "fmt"

// ServerHacker is *SOMETHING* that can hack servers.  We don't care what it is,
// how it came to be, or anything like that.  We only care that it has a receiver
// named HackServers() that takes no arguments and return an error.
type ServerHacker interface {
	HackServers() error
}

// Server is our API that lets us trigger useful behavior, like hacking servers
type Server struct {
	hacker ServerHacker
}

// New returns a new server ready to listen
func New(hacker ServerHacker) *Server {
	return &Server{
		hacker: hacker,
	}
}

// HackStuff would be some HTTP handler, but for simplicity we'll just make it
// callable here.  Hack stuff using our ServerHacker we got on creation.
func (s *Server) HackStuff() error {
	fmt.Println("About to hack stuff from the server...")
	return s.hacker.HackServers()
}
