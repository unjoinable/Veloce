package main

import (
	"Veloce/internal/network/server"
	"log"
	"runtime"
	"time"
)

func main() {
	go monitorResources()
	startServer()
}

func startServer() {
	srv := server.NewMinecraftServer()
	srv.Init()
	srv.Start("127.0.0.1:25565")
}

func monitorResources() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		log.Printf("=== Resource Usage ===")
		log.Printf("Memory: %.2f MB allocated | %.2f MB system",
			float64(m.Alloc)/1024/1024,
			float64(m.Sys)/1024/1024)
		log.Printf("Goroutines: %d | GC Runs: %d | Next GC: %.2f MB",
			runtime.NumGoroutine(),
			m.NumGC,
			float64(m.NextGC)/1024/1024)
	}
}
