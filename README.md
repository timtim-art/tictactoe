# Broadcasting Prototype

start server (install [go first](https://go.dev/doc/install)):
```sh
cd backend-go
go run cmd/main.go
```

setup game (game id will be returned)
```sh
curl "localhost:8080/api/v1/fight/setup" 
```

start serving react app
```sh
cd frontend
npm run dev
```

call `http://localhost:5173/` in browser (try firefox or chrome if it does not work) in multiple tabs

start fight
```sh
curl "localhost:8080/api/v1/fight/start?gameId=1"
```

# Live Streaming Data Research

Prerequisites:
- Client is a browser
- Latency is key

| Technology | Browser Support | Protocol underneath | Latency | Reliability | Implementation Complexity | Use Case |
| ---------- | --------- | --------- | --------- | --------- | --------- | --------- | 
| WebSockets | Yes | TCP | High | High | Easy | Messengers like WhatsApp |
| WebTransport | Yes, but not Safari | QUIC (over UDP) | Middle | Very High | Middle | New kid in the hood |
| WebRTC | Yes | UDP (TCP fallback) | Low | Low | Hard* | Video conference tools like Google Meet or MS Teams |
| Streaming over HTTP | Yes | TCP | Very High | High | Middle | Video Live Streaming (IPTV) |

\*separate server (STUN/TURN) for establishing connection is needed 