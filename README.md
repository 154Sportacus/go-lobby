# go-lobby
> in-memory shared data-exchange session
 
[![Go Reference](https://pkg.go.dev/badge/github.com/troublete/go-lobby.svg)](https://pkg.go.dev/github.com/troublete/go-lobby)

## Introduction

This library provides a serve mux for transient shared data-exchange sessions between multiple clients; As the demo
shows it can be used to build a rudimentary chat server or a custom server for exchanging data between multiple
clients without the need for a database*. For a proper implementation the server should define filter methods for user
connections and sent messages (like format validation or checks for offensive content and the like), but the raw default
implementation with only client side logic works very well also, as the demos showcase.

\* this means that servers created with this library do not scale to more than one instance behind a load balancer, as
any lobby created is very much server bound as it is stored in-memory only

## Quickstart

```
package main

import (
	"log"
	"net/http"

	"github.com/troublete/go-lobby/server"
)

func main() {
	err := http.ListenAndServe(":1234", server.NewLobbyMux())
	if err != nil {
		log.Fatal(err)
	}
}
```

This snippet will create a server with two routes exposed

1. `GET /connect/{lobbyIdx}/{clientIdx}?captionName={human readable clientName}` (`captionName` is optional) 

This is an SSE (Server-sent events) endpoint, which is used to connect to a shared lobby; if the lobby doesn't exist
it'll spring into existence upon creation. The default expiration of a lobby is set to 5 minutes. A client can only
connect if the quotas for lobbies (as the creation is implicit) or for clients in a lobby isn't exceeded.

Aside the expiration a lobby also gets cleaned up after the last client "leaves".

Upon connection the client immediately receives to messages; one with the first token (which is necessary as
authentication for sending messages back) and another one with the expiration time of the session. The token is after
each send message renewed by sending the client a similar message as the first one with a new token.

All other messages that the client receives are user defined. So you can use this server freely for multiple use-cases
without needing to change the server.

For further customisation a sole argument can be passed to the `NewLobbyMux` which is a configuration which allows for
different quotas and adding filter functions for more advanced uses which require user or message validation or
filtering.

Server-sent events can be easily consumed on the javascript side through the `EventSource` object. Following a quick
example.

```
const es = new EventSource("/connect/...");
es.onmessage = (e) => {
    // use e.data
};
```

It is also advisable to use as a client id an uuid or similar, same for lobbies (which is also available through a
baseline browser functionality), just to be sure that a session is unique and only meant-to-be-in clients join.

2. `POST /send/{lobbyIdx}/{clientIdx}` (requires `Authorization` header with a `Token *` value)

This endpoint provides the functionality to send data in a broadcasting fashion to all connected clients of a lobby. It
requires a per-request authentication as described above. The tokens are received through the SSE-endpoint and are
unique to each client.

The data is sent through the body and its meant to be JSON. The default configuration limits the message size sent to 100
bytes. This can be changed through the config if needed.

Similar to the connection endpoint the send endpoint also allows a filter configuration through the config, this might
be relevant to validate structure of a message or do actual content validation.

## License

All rights reserved