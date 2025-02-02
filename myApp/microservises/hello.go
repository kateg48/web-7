package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getMessage", handleGetRequest)
	fmt.Println("Starting server on port 8081")
	http.ListenAndServe(":1232", nil)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "Hello, web!")
}
