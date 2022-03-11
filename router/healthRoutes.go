package router

import (
	"github.com/gin-gonic/gin"
)

func healthShow(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
