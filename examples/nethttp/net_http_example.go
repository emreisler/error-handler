package main

import (
	"github.com/emreisler/error-handler"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusBadRequest)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/example", handler)

	wrappedMux := error_handler.NetHTTPMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
