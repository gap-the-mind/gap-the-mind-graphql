// +build !js

package main

import (
	"os"

	"github.com/gap-the-mind/gap-the-mind-graphql/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server.StartServer(port)
}
