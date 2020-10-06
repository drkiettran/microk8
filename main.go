package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {fmt.Fprint(w, "hello from web page")})
	http.ListenAndServe(":8080", nil)
}
