package websocket

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Viewer struct {
	conn *websocket.Conn
	hub  *Hub
	send chan []byte
	id   int
}

func NewViewer(hub *Hub, c *gin.Context) (*Viewer, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error %s when upgrading to websocket", err)
		return nil, err
	}
	Viewer := &Viewer{hub: hub, conn: conn, send: make(chan []byte, 256)}
	Viewer.hub.register <- Viewer

	go Viewer.write()

	return Viewer, nil
}

func (p *Viewer) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		p.conn.Close()
	}()
	for {
		select {
		case message := <-p.send:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := p.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Error %s when getting writer in viewer", err)
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				log.Printf("Error %s when closing writer in viewer", err)
				return
			}
		case <-ticker.C:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := p.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("Error %s when pinging peer", err)
				return
			}
		}
	}
}
