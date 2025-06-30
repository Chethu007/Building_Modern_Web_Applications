package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func main() {
	http.HandleFunc("/", Home)
	_ = http.ListenAndServe(port, nil)
}
