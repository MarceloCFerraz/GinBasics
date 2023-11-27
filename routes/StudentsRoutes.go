package AlunosRoutes

import (
	"github.com/MarceloCFerraz/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	server := gin.Default()
	
	server.GET("/students", controllers.)

	server.Run() // defaults to port 8080, provide ":5050" to change to another port
}