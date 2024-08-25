# Live Streaming Data

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