package main

import (
	"fmt"
	"net/http"

	"github.com/Fepelus/clockgsl"
)

const port = ":7799"

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/string", stringhandler)
	fmt.Printf("http://localhost%s/\n", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, clockgsl.GetAllAsJSON())
}

func stringhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, clockgsl.GetAllAsString())
}
