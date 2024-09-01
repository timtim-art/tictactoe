package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/timtim-art/tictactoe/backend-go/internal/server/websocket"
)

const API_VERSION string = "/api/v1"

func New() *gin.Engine {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	gameHandler := websocket.NewGameHandler()
	router.GET(API_VERSION+"/fight/setup", gameHandler.SetupFight)
	router.GET(API_VERSION+"/fight/start", gameHandler.StartFight)
	router.GET(API_VERSION+"/stream", gameHandler.GetStreams)
	router.GET(API_VERSION+"/stream/join", gameHandler.JoinStream)

	return router
}
