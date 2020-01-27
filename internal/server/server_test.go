package server

import "testing"

// mockServerHacker pretends to hack things.  It doesn't actually do anything,
// but it will track whether or not it was called.
type mockServerHacker struct {
	called bool
}

// Implement our ServerHacker interface... note we're not mocking everything
// that is in the concrete package, only what's in the interface for server.
func (h *mockServerHacker) HackServers() error {
	// Track that we were called
	h.called = true
	return nil
}

func newMockServerHacker() *mockServerHacker {
	return &mockServerHacker{
		// This isn't necessary, but yay being explicit
		called: false,
	}
}

func TestHackStuffCallsUnderlyingClient(t *testing.T) {
	mockClient := newMockServerHacker()

	// Note that we can use the mock client here just the same as we did for
	// the main server because we've implemented the interface fully
	s := New(mockClient)

	s.HackStuff()

	if !mockClient.called {
		t.Error("Expected client to be called, but was not")
	}
}
