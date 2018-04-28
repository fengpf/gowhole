package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func showIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ShowIndex!\n"))
}

func showAdminDashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showAdminDashboard!\n"))
}

func main() {
	r := http.NewServeMux()

	// Only log requests to our admin dashboard to stdout
	r.Handle("/admin", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(showAdminDashboard)))
	r.HandleFunc("/", showIndex)

	// Wrap our server with our gzip handler to gzip compress all responses.
	http.ListenAndServe(":8000", handlers.CompressHandler(r))
}
