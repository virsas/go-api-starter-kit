package example

import (
	"encoding/json"
	"go-api-starter-kit/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (env *env) list(ctx *gin.Context) {
	s := &service{db: env.db, log: env.log, ctx: env.ctx}
	examples, err := s.list()
	if err != nil {
		ctx.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	ctx.JSON(config.OK_STATUS, examples)
}

func (env *env) create(ctx *gin.Context) {
	var example examplereq
	err := json.NewDecoder(ctx.Request.Body).Decode(&example)
	if err != nil {
		env.log.Error(config.SERVER_STRING, zap.Error(err))
		ctx.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	s := &service{db: env.db, log: env.log, ctx: env.ctx}
	err = s.create(example)
	if err != nil {
		ctx.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	ctx.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}

func (env *env) show(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
	if err != nil {
		env.log.Error(config.SERVER_STRING, zap.Error(err))
		ctx.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	s := &service{db: env.db, log: env.log, ctx: env.ctx}
	example, err := s.show(id)
	if err != nil {
		ctx.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	ctx.JSON(config.OK_STATUS, example)
}

func (env *env) update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
	if err != nil {
		env.log.Error(config.SERVER_STRING, zap.Error(err))
		ctx.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	var example examplereq
	err = json.NewDecoder(ctx.Request.Body).Decode(&example)
	if err != nil {
		env.log.Error(config.SERVER_STRING, zap.Error(err))
		ctx.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	s := &service{db: env.db, log: env.log, ctx: env.ctx}
	err = s.update(id, example)
	if err != nil {
		ctx.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	ctx.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}

func (env *env) delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("ID"), 10, 64)
	if err != nil {
		env.log.Error(config.SERVER_STRING, zap.Error(err))
		ctx.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	s := &service{db: env.db, log: env.log, ctx: env.ctx}
	err = s.delete(id)
	if err != nil {
		ctx.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	ctx.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}
