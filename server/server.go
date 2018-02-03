package main

import (
	"net/http"
	"fmt"
)

func ping(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "pong")
}

func pong(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "ping")
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/pong", pong)
	http.ListenAndServe(":8000", nil)
}
