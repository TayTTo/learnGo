package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var pong string = "pong"

func getPing(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, pong)
}
func main() {
	r := gin.Default()
	r.GET("/ping", getPing)
	r.Run("localhost:8081")
}
