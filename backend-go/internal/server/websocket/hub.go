package websocket

// Hub maintains the set of active Viewers and broadcasts messages to the
// Viewers.
type Hub struct {
	// Registered Viewers.
	Viewers map[int]*Viewer

	// Inbound messages from the Viewers.
	broadcast chan []byte

	// Register requests from the Viewers.
	register chan *Viewer

	// Unregister requests from Viewers.
	unregister chan *Viewer

	counter int
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Viewer),
		unregister: make(chan *Viewer),
		Viewers:    make(map[int]*Viewer),
		counter:    0,
	}
}

func (h *Hub) run() {
	for {
		select {
		case Viewer := <-h.register:
			h.counter++
			Viewer.id = h.counter
			h.Viewers[h.counter] = Viewer
		case Viewer := <-h.unregister:
			if _, ok := h.Viewers[Viewer.id]; ok {
				delete(h.Viewers, Viewer.id)
				close(Viewer.send)
			}
		case message := <-h.broadcast:
			for id, Viewer := range h.Viewers {
				select {
				case Viewer.send <- message:
				default:
					close(Viewer.send)
					delete(h.Viewers, id)
				}
			}
		}
	}
}
