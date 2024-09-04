package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/timtim-art/tictactoe/backend-go/internal/gameplay"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type GameHandler struct {
	Games   map[int]*Hub
	Counter int
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		Games:   make(map[int]*Hub),
		Counter: 0,
	}
}

func (g *GameHandler) SetupFight(c *gin.Context) {
	g.Counter++
	hub := newHub()
	g.Games[g.Counter] = hub
	go hub.run()

	c.JSON(http.StatusAccepted, gin.H{
		"gameId": g.Counter,
	})
}

func (g *GameHandler) StartFight(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Query("gameId"))
	if err != nil {
		log.Printf("Error %s when converting game id", err)
		c.Status(http.StatusBadRequest)
		return
	}

	game, ok := g.Games[gameId]
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	go sendGame(game)
}

func sendGame(game *Hub) {
	warriors := gameplay.CreateWarriorsRandomly()
	for {
		if gameplay.Finished(warriors) {
			return
		}

		b, err := json.Marshal(warriors)
		if err != nil {
			log.Printf("Error %s when marshalling warriors", err)
			return
		}
		game.broadcast <- b

		gameplay.CalcPositions(warriors)

		time.Sleep(40 * time.Millisecond)
	}
}

func (g *GameHandler) GetStreams(c *gin.Context) {
	keys := make([]int, len(g.Games))
	i := 0
	for k := range g.Games {
		keys[i] = k
		i++
	}
	c.JSON(http.StatusAccepted, gin.H{
		"games": keys,
	})
}

func (g *GameHandler) JoinStream(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Query("gameId"))
	if err != nil {
		log.Printf("Error %s when converting game id", err)
		c.Status(http.StatusBadRequest)
		return
	}

	game, ok := g.Games[gameId]
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err = NewViewer(game, c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}
