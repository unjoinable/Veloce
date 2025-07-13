# Running Veloce

Veloce is a lightweight Minecraft server written in Go. Currently, it only supports basic server ping and MOTD functionality.
Prerequisites

    Go 1.20 or newer installed

    A Minecraft client (1.8+ recommended for testing)

# Getting Started

To use Veloce, simply create a new Go file and include the following code:

package main

import "github.com/your-username/veloce/server" // Replace with actual module path

func main() {
    srv := server.NewMinecraftServer()
    srv.Init()
    srv.Start("127.0.0.1:25565")
}

This will start the server on 127.0.0.1:25565, allowing your Minecraft client to connect and receive the MOTD response.
Notes

    Only server ping (MOTD) is supported.

    No gameplay or world functionality is implemented (yet).

    This project is a minimal, educational base for building Minecraft server features in Go.
