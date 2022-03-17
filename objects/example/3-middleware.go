package example

import (
	"bytes"
	"encoding/json"
	"go-api-starter-kit/config"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func validateExample(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		var example examplereq
		err = json.Unmarshal(data, &example)
		if err != nil {
			c.JSON(config.REQUEST_ERROR, gin.H{
				"message": config.REQUEST_STRING,
			})
			c.Abort()
			return
		}

		validate := validator.New()
		err = validate.Struct(&example)
		if err != nil {
			log.Error("Validation error", zap.Error(err))
			c.JSON(config.VALID_ERROR, gin.H{
				"message": config.VALID_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
