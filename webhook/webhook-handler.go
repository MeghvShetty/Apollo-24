package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebhookPayload represents the JSON structure of the webhook payload
type WebhookPayload struct {
	Issue Issue `json:"issue"`
}

// Issue represents the 'issue' field in the webhook payload
type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

// Fields represents the 'fields' field within the Issue in the webhook payload
type Fields struct {
	Summary       string   `json:"summary"`
	Assignee      Assignee `json:"assignee"`
	CustomField34 string   `json:"customfield_10034"`
	Description   string   `json:"description"`
	Labels        []string `json:"labels"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

// WebhookHandler handles incoming webhook requests
func WebhookHandler(c *gin.Context) {
	var payload WebhookPayload

	// Bind JSON payload to WebhookPayload struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Print the received payload for debugging
	fmt.Println("Received Webhook Payload:")
	prettyPrintJSON(payload)

	// Process the webhook payload as needed (e.g., save to database, trigger actions, etc.)

	// Respond to the webhook request
	c.JSON(http.StatusOK, gin.H{
		"message": "Webhook received and processed successfully",
	})
}

// prettyPrintJSON prints JSON in an indented format
func prettyPrintJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(b))
}
