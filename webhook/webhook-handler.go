package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookPayload struct {
	Timestamp          int64     `json:"timestamp"`
	WebhookEvent       string    `json:"webhookEvent"`
	IssueEventTypeName string    `json:"issue_event_type_name"`
	User               User      `json:"user"`
	Issue              Issue     `json:"issue"`
	Changelog          Changelog `json:"changelog"`
}

type User struct {
	Self      string `json:"self"`
	AccountID string `json:"accountId"`
	// AvatarUrls  AvatarUrls `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	AccountType string `json:"accountType"`
}

// type AvatarUrls struct {
// 	Size48x48 string `json:"48x48"`
// 	Size24x24 string `json:"24x24"`
// 	Size16x16 string `json:"16x16"`
// 	Size32x32 string `json:"32x32"`
// }

type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	StatusCategoryChangeDate string      `json:"statuscategorychangedate"`
	IssueType                IssueType   `json:"issuetype"`
	TimeSpent                interface{} `json:"timespent"`
	// Project                  Project     `json:"project"`
	CustomField10034   string      `json:"customfield_10034"`
	AggregateTimeSpent interface{} `json:"aggregatetimespent"`
	Resolution         Resolution  `json:"resolution"`
	ResolutionDate     string      `json:"resolutiondate"`
	WorkRatio          int         `json:"workratio"`
	LastViewed         string      `json:"lastViewed"`
	// Watches            Watches     `json:"watches"`
	Created string `json:"created"`
	// Priority           Priority    `json:"priority"`
	Labels      []string `json:"labels"`
	Assignee    Assignee `json:"assignee"`
	Updated     string   `json:"updated"`
	Status      Status   `json:"status"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Creator     User     `json:"creator"`
	Reporter    User     `json:"reporter"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}

type Project struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	ProjectTypeKey string `json:"projectTypeKey"`
	Simplified     bool   `json:"simplified"`
	// AvatarUrls     AvatarUrls `json:"avatarUrls"`
}

type Resolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// type Watches struct {
// 	Self       string `json:"self"`
// 	WatchCount int    `json:"watchCount"`
// 	IsWatching bool   `json:"isWatching"`
// }

// type Priority struct {
// 	Self    string `json:"self"`
// 	IconUrl string `json:"iconUrl"`
// 	Name    string `json:"name"`
// 	ID      string `json:"id"`
// }

type Assignee struct {
	Self      string `json:"self"`
	AccountID string `json:"accountId"`
	// AvatarUrls  AvatarUrls `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	AccountType string `json:"accountType"`
}

type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconUrl        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type Changelog struct {
	ID    string       `json:"id"`
	Items []ChangeItem `json:"items"`
}

type ChangeItem struct {
	Field      string      `json:"field"`
	FieldType  string      `json:"fieldtype"`
	FieldID    string      `json:"fieldId"`
	From       interface{} `json:"from"`
	FromString interface{} `json:"fromString"`
	To         interface{} `json:"to"`
	ToString   string      `json:"toString"`
}

func WebhookHandler(c *gin.Context) {
	var payload WebhookPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Print the received payload for debugging
	fmt.Println("Received Webhook Payload:")
	prettyPrintJSON(payload)

	// Process the webhook payload as needed (e.g., save to database, trigger actions, etc.)

	// NewIssue := &jira.AARep{
	// 	IssueName:    payload.Issue.Fields.Summary,
	// 	AssigneeName: payload.Issue.Fields.Assignee.DisplayName,
	// 	DueDate:      payload.Issue.Fields.CustomField10034,
	// 	Description:  payload.Issue.Fields.Description,
	// 	SIRrating:    payload.Issue.Fields.Labels,
	// 	Changelog:    payload.Issue.Fields.Status.StatusCategory.Key,
	// }

	// // jira.CreateIssueAA(NewIssue)
	// fmt.Println(NewIssue)

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received and processed successfully"})
}

func prettyPrintJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(b))
}
