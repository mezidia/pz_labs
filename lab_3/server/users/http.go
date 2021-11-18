package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mezidia/pz_labs/lab_3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListUsers(store, rw)
		} else if r.Method == "POST" {
			handleUserCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateUser(u.UserName)
	if err == nil {
		tools.WriteJsonOk(rw, &u)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListUsers(store *Store, rw http.ResponseWriter) {
	res, err := store.ListUsers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
