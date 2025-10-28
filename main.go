package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"html/*"
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
	reqFile := r.URL.Path
	if reqFile == "/" {
		reqFile += "index.htjs"
	}
	fmt.Printf("GET:  %s\n", reqFile)
	file, err := os.ReadFile(reqFile[1:])
	if err != nil {
		fmt.Println(err)
	}

	

	w.Write(file)
}
