package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Save the request body to a file
	saveEventToFile(body)

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received and saved webhook event"))
}

func saveEventToFile(data []byte) {
	f, err := os.OpenFile("webhook_events.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer f.Close()

	// Append timestamp to the payload
	timestamp := time.Now().Format(time.RFC3339)
	payload := timestamp + "\n" + string(data) + "\n\n"

	// Write payload to file
	if _, err := f.WriteString(payload); err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
