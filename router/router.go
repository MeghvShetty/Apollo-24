/*
Apollo-24 routing handler
*/

package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebServer() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	log.Fatal(router.Run(":3000"))
}
