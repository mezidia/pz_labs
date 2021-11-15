package main

import (
	"log"
	"os"

	"github.com/mezidia/pz_labs/lab_3/server/channels"
)

func ComposeApiServer(port HttpPortNumber) *ChatApiServer {
	chatApiServer := &ChatApiServer{
		Port:   port,
		router: ComposeRouter(),
	}
	return chatApiServer
}

func ComposeForumsHandler() channels.HttpHandlerFunc {
	db, err := NewDbConnection()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := channels.NewStore(db)
	httpHandlerFunc := channels.HttpHandler(store)
	return httpHandlerFunc
}

// ComposeRouter will create an instance of Router according to providers defined in this file.
func ComposeRouter() *Router {
	router := &Router{ROUTER_CONFIG}
	return router
}
