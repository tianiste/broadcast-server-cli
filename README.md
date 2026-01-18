# Broadcast Server CLI

A simple broadcast chat server built with websockets in go

## Features

- Multiple clients supported at the same time
- Messages are broadcast to all connected clients
- Terminal (CLI) client
- Username provided via URL path


## Requirements

- Go 1.20+ (earlier versions may also work)
- Gorilla WebSocket library

Install dependencies:

```bash
go get github.com/gorilla/websocket
```


## How It Works

- The server exposes a WebSocket endpoint at `/ws/{username}`
- Each client connects using a username
- When a client sends a message:
  - the server receives it
  - the server broadcasts it to all connected clients
- The server keeps track of connected clients in memory


## Running the Server

Start the server on a chosen port:

```bash
go run main.go broadcast-server -p <port> -s
```
## Connecting as a Client

Open a new terminal for each client and run:

```bash
go run . broadcast-server  -p <port> -u <username> -c
```
Typing a message in one terminal will broadcast it to all connected clients.

## Notes

- There is no authentication or persistence
- Messages are not stored
- Connections are not encrypted (no TLS)

## Possible Improvements

- Per-client write goroutines
- Message history
- Authentication
- Structured JSON messages

[Project Idea](https://roadmap.sh/projects/broadcast-server)
