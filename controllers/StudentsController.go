package StudentsController

import (
	"github.com/gin-gonic/gin"
)

func GetStudents(ctx *gin.Context) {
	ctx.JSON(
		200,
		gin.H{
			"id": "1",
			"name": "Marcelo",
		},
	)
}