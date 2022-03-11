package router

import (
	"encoding/json"
	"go-api-starter-kit/controllers"
	"go-api-starter-kit/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (env *Env) exampleList(ctx *gin.Context) {
	status, message, result := controllers.ExampleList(env.db, env.log, env.ctx)
	ctx.JSON(status, gin.H{
		"message": message,
		"payload": result,
	})
}

func (env *Env) exampleCreate(ctx *gin.Context) {
	var model models.ExampleReq
	err := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if err != nil {
		env.log.Error("Error", zap.Error(err))
		ctx.JSON(400, gin.H{
			"message": "apiIssue",
		})
		return
	}

	status, message := controllers.ExampleCreate(model, env.db, env.log, env.ctx)
	ctx.JSON(status, gin.H{
		"message": message,
	})
}

func (env *Env) exampleShow(ctx *gin.Context) {
	idParam := ctx.Param("ID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		env.log.Error("Error", zap.Error(err))
		ctx.JSON(400, gin.H{
			"message": "apiIssue",
		})
		return
	}
	status, message, result := controllers.ExampleShow(id, env.db, env.log, env.ctx)
	ctx.JSON(status, gin.H{
		"message": message,
		"payload": result,
	})
}

func (env *Env) exampleUpdate(ctx *gin.Context) {
	idParam := ctx.Param("ID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		env.log.Error("Error", zap.Error(err))
		ctx.JSON(400, gin.H{
			"message": "apiIssue",
		})
		return
	}

	var model models.ExampleReq
	err = json.NewDecoder(ctx.Request.Body).Decode(&model)
	if err != nil {
		env.log.Error("Error", zap.Error(err))
		ctx.JSON(400, gin.H{
			"message": "apiIssue",
		})
		return
	}

	status, message := controllers.ExampleUpdate(id, model, env.db, env.log, env.ctx)
	ctx.JSON(status, gin.H{
		"message": message,
	})
}

func (env *Env) exampleDelete(ctx *gin.Context) {
	idParam := ctx.Param("ID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		env.log.Error("Error", zap.Error(err))
		ctx.JSON(400, gin.H{
			"message": "apiIssue",
		})
		return
	}
	status, message := controllers.ExampleDelete(id, env.db, env.log, env.ctx)
	ctx.JSON(status, gin.H{
		"message": message,
	})
}
