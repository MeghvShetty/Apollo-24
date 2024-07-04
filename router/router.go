/*
Apollo-24 routing handler
*/

package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func WebServer() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("static files/index.html")
	})
	log.Fatal(router.Run(":3000"))
}
