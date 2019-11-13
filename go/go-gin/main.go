package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World") //  send status --> 200 and string --> "hello world"
	})
	r.Run() //  listen on port :8080
}
