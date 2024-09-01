package main

import "github.com/timtim-art/tictactoe/backend-go/internal/server"

func main() {
	server := server.New()
	server.Run()
}
