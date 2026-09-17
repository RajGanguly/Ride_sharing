package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(HandleTripPreview))
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
