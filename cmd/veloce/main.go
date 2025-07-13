package main

import (
	"Veloce/internal/network/server"
)

func main() {
	startServer()
}

func startServer() {
	srv := server.NewMinecraftServer()
	srv.Init()
	srv.Start("127.0.0.1:25565")
}
