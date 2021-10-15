package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func handleReq(res http.ResponseWriter, req *http.Request) {
	rfc_time := time.Now().Format(time.RFC3339)
	res_json, err := json.Marshal(rfc_time)
	if err != nil {
		fmt.Fprintf(res, "Error json marshal\n")
	}
	fmt.Fprintf(res, "Time now: %s\n", res_json)

}

func main() {

	http.HandleFunc("/", handleReq)

	http.ListenAndServe(":8080", nil)

	fmt.Println("server run at http://localhost:8080/")

}
