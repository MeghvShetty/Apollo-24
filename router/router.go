/*
Apollo-24 routing handler
*/

package router

import (
	"fmt"
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
		fmt.Println(logic.SirCalculator(&triageInput))
	})

	// router.POST("/submit-form", func(ctx *gin.Context) {
	// 	var triageInput logic.TriageInput
	// 	if err := ctx.ShouldBind(&triageInput); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// })
	log.Fatal(router.Run(":3000"))
}
