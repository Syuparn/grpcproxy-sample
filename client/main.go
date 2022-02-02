package main

import (
	"fmt"
	"net/http"
)

// run dummy server only for sidecar health check
func dummyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dummy")
}

func main() {
	// request to hello server
	go request()

	http.HandleFunc("/", dummyHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
