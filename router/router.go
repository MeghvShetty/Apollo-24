/*
Apollo-24 routing handler
*/

package router

import (
	"log"
	"net/http"

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
		if err := ctx.ShouldBind(&triageInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		output := logic.SirCalculator(&triageInput)

		ctx.HTML(http.StatusOK, "response.html", gin.H{
			"output":      output,
			"TriageInput": triageInput,
		})
	})

	log.Fatal(router.Run(":3000"))
}
