package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type response struct {
	Time string `json:"time"`
}

func main() {

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		result, _ := json.Marshal(response{time.Now().Format(time.RFC3339)})
		fmt.Fprintf(w, string(result))
	})

	fmt.Printf("Starting server at port 8795\n")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}
