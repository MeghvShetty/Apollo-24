/*
Apollo-24 routing handler
*/

package router

import (
	"log"
	"net/http"

	"github.com/The-Manchester-Project/Apollo-24/api/jira"
	"github.com/The-Manchester-Project/Apollo-24/logic"
	"github.com/gin-gonic/gin"
)

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

		ctx.HTML(http.StatusOK, "response.html", gin.H{
			"SirRating":     sir_rating,
			"TriageInput":   triageInput,
			"IssueTemplate": jiraTemple,
		})
	})

	log.Fatal(router.Run(":3000"))
}
