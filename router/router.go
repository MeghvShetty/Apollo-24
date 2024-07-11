/*
Apollo-24 routing handler
*/

package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/The-Manchester-Project/Apollo-24/api/jira"
	"github.com/The-Manchester-Project/Apollo-24/logic"
	"github.com/gin-gonic/gin"
)

type CreateIssueResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

func WebServer() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/submit-form", func(ctx *gin.Context) {
		var triageInput logic.TriageInput
		var jiraTemple jira.IssueTemplate
		if err := ctx.ShouldBind(&triageInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctx.ShouldBind(&jiraTemple); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sir_rating := logic.SirCalculator(&triageInput)
		jiraTemple.SIRrating = []string{sir_rating}
		createTemple, err := jira.CreateIssue(&jiraTemple)
		if err != nil {
			fmt.Errorf("failed to fetch Jira metadata: %w", err)
		}

		var createIssueResponse CreateIssueResponse
		if err := json.Unmarshal([]byte(createTemple), &createIssueResponse); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse Jira response"})
			return
		}

		jiraKey := createIssueResponse.Key

		ctx.HTML(http.StatusOK, "response.html", gin.H{
			"SirRating":     sir_rating,
			"TriageInput":   triageInput,
			"IssueTemplate": jiraTemple,
			"Jira":          jiraKey,
		})
	})

	log.Fatal(router.Run(":3000"))
}
