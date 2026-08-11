package main

import (
	"fmt"
	"net/http"
)

// Handler is the Vercel entrypoint for /api/hello when using /api/$1.go route mapping.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from /api/hello")
}
