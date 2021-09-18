package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "time!")
    })

    fmt.Printf("Starting server at port 8795\n")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        log.Fatal(err)
    }
}