package main

import (
	"log"
	"os"

	"github.com/mezidia/pz_labs/lab_3/server/forums"
	"github.com/mezidia/pz_labs/lab_3/server/users"
)

func ComposeApiServer(port HttpPortNumber) *ApiServer {
	chatApiServer := &ApiServer{
		Port:   port,
		router: ComposeRouter(),
	}
	return chatApiServer
}

func ComposeForumsHandler() forums.HttpHandlerFunc {
	db, err := NewDbConnection()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := forums.NewStore(db)
	httpHandlerFunc := forums.HttpHandler(store)
	return httpHandlerFunc
}

func ComposeUsersHandler() users.HttpHandlerFunc {
	db, err := NewDbConnection()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := users.NewStore(db)
	httpHandlerFunc := users.HttpHandler(store)
	return httpHandlerFunc
}

// ComposeRouter will create an instance of Router according to providers defined in this file.
func ComposeRouter() *Router {
	router := &Router{ROUTER_CONFIG}
	return router
}
