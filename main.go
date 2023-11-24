package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Run() // defaults to port 8080, provide ":5050" to change to another port
}