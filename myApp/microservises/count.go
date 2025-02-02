package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var cnt int = 0

func handleCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%d", cnt)
	case "POST":
		count, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		cnt += count
	default:
		http.Error(w, "Неизвестный метод", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/count", handleCount)
	fmt.Println("Starting server on port 8083")
	http.ListenAndServe(":8083", nil)
}
