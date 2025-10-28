package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

const (
	port = 8080
)

func main() {
	fmt.Printf("listening on port: %d\n", port)

	http.HandleFunc("/", httpServer)

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func httpServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET:  %s\n", r.URL.Path)
	file, err := os.ReadFile(r.URL.Path[1:])
	if err != nil {
		fmt.Println(err)
	}
	w.Write(file)
}
