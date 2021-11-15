package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mezidia/pz_labs/tree/lab3/lab_3/server/db"
)

var httpPortNumber = flag.Int("p", 8081, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "chat-example",
		User:       "roman",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	// Parse command line arguments. Port number may be defined with "-p" flag.
	flag.Parse()
	// Create the server.
	server := ComposeApiServer(HttpPortNumber(*httpPortNumber))
	// Start it.
	go func() {
		log.Println("Starting chat server...")

		err := server.Start()
		if err == http.ErrServerClosed {
			log.Printf("HTTP server stopped")
		} else {
			log.Fatalf("Cannot start HTTP server: %s", err)
		}
	}()

	// Wait for Ctrl-C signal.
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel

	if err := server.Stop(); err != nil && err != http.ErrServerClosed {
		log.Printf("Error stopping the server: %s", err)
	}
}
