package main

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol"
	"log"
	"runtime"
	"time"
)

func main() {
	protocol.RegisterAllPackets()

	// Start resource monitoring in a separate goroutine
	go monitorResources()

	// Create the server
	srv := network.NewTCPServer("127.0.0.1:25565", 100) // 100 max connections

	// Start the server and block
	if err := srv.Start(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
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
