package middlewares

import (
	"database/sql"
	"go-api-starter-kit/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type usermodel struct {
	aid    sql.NullInt64
	locked sql.NullBool
}

func User(db *sql.DB, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var user usermodel

		uid := c.MustGet("uid")

		err = db.QueryRow("SELECT account_id, locked FROM users WHERE id=?;", uid).Scan(&user.aid, &user.locked)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Error("notFound", zap.Error(err))
				c.JSON(config.NOTFOUND_ERROR, gin.H{
					"message": config.NOTFOUND_STRING,
				})
				c.Abort()
				return
			}
			log.Error("dbError", zap.Error(err))
			c.JSON(config.DB_ERROR, gin.H{
				"message": config.DB_STRING,
			})
			c.Abort()
			return
		}

		if user.locked.Valid && user.locked.Bool {
			c.JSON(config.AUTH_ERROR, gin.H{
				"message": config.AUTH_STRING,
			})
			c.Abort()
			return
		}

		if user.aid.Valid {
			c.Set("aid", user.aid.Int64)
		} else {
			log.Error("notFound", zap.Error(err))
			c.JSON(config.NOTFOUND_ERROR, gin.H{
				"message": config.NOTFOUND_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
