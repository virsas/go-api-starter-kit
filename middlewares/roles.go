package middlewares

import (
	"go-api-starter-kit/config"
	"go-api-starter-kit/utils/logger"

	"github.com/gin-gonic/gin"
)

func AllowRoles(log *logger.Logger, allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesString := c.MustGet("roles")
		userRoles := make([]string, len(rolesString.([]interface{})))
		for i, v := range rolesString.([]interface{}) {
			userRoles[i] = v.(string)
		}

		if !includes(userRoles, allowedRoles) {
			log.Error("Not allowed to access")
			c.JSON(config.AUTH_ERROR, gin.H{
				"message": config.AUTH_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// includes function
func includes(userRoles []string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		for _, userRole := range userRoles {
			if allowedRole == userRole {
				return true
			}
		}
	}
	return false
}
