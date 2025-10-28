package main

import (
	"fmt"
	"strconv"
	"net/http"
)

const (
	port = 8080
)

func main() {
	fmt.Printf("listening on port: %d", port)

	http.HandleFunc("/", httpServer)

	panic(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func httpServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)
}
