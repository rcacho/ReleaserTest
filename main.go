package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusOK)
	})

	engine.GET("/version", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusAccepted)
	})

	if err := engine.Run(); err != nil {
		log.Println(err)
	}
}
