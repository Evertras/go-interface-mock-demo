# go-interface-mock-demo

Quick demonstration of how to use Go interfaces to mock existing concrete libraries.

There is an example concrete package in [pkg/external/concrete](pkg/external/concrete).
It contains some client that does two useful things: hack servers and mine bitcoins.

Our server that we're writing for this application is interested in hacking servers, so
we've imported this concrete package to do that for us.  Unfortunately the concrete package
has not supplied us with an interface; we can ONLY get a concrete implementation out of it.

The problem is we still want to write tests for our code.  So how do we test with a client
that will reach external resources if we let it?

Go interfaces to the rescue!

Go interfaces aren't like other languages where a class must inherit the interface
in order to match the interface type later.  Go takes the duck typing approach
of "If this struct can do everything this interface asks it to, then it must be
this interface."

Using this knowledge, we can actually remove the `concrete.ConcreteClient` type
from `server` entirely.  Instead, `server` will describe an interface of a client
that it needs in order to function.  This interface is designed to match the
methods that we're interested in from `concrete.ConcreteClient`.  Now when we
create a new server, we can pass in a `*concrete.ConcreteClient` instance because
it matches our interface.  We can also create a mock implementation very easily
for our unit tests.  Check out [the unit tests here.](./internal/server/server_test.go).

