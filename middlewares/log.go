package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-starter-kit/utils"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type logBody struct {
	User     string `json:"user"`
	Action   string `json:"action"`
	Resource string `json:"resource"`
	Body     string `json:"body"`
}
type logMessage struct {
	Message logBody `json:"message"`
	Level   string  `json:"level"`
	Label   string  `json:"label"`
}

func Log(audit *utils.Audit, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		email := fmt.Sprintf("%v", c.MustGet("email"))
		body := &logBody{User: email, Action: c.Request.Method, Resource: c.Request.RequestURI, Body: string(data)}
		message := &logMessage{Message: *body, Level: "info", Label: "audit"}
		finalMessage, err := json.Marshal(message)
		if err != nil {
			log.Error("JSON marshal error", zap.Error(err))
		}

		err = utils.CwWriteLog(audit, string(finalMessage))
		if err != nil {
			log.Error("JSON marshal error", zap.Error(err))
		}
		c.Next()
	}
}
