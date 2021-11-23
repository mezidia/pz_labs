package interests

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
			handleListInterests(store, rw)
		} else if r.Method == "POST" {
			handleInterestCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleInterestCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var i Interest
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Printf("Error decoding interest input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateInterest(i.InterestName, i.ForumId)
	if err == nil {
		tools.WriteJsonOk(rw, &i)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListInterests(store *Store, rw http.ResponseWriter) {
	res, err := store.ListInterests()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
