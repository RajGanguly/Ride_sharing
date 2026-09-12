package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	log.Println("Received request:", reqBody)
	log.Println("JSON Encoder", r.Body)
	if reqBody.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
	}

	log.Println("SUCCESS")

	response := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, response)
}
