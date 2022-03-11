package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Example() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.Header.Get("User-Agent"))
	}
}
