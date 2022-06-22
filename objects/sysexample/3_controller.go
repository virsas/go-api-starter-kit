package sysexample

import (
	"database/sql"
	"encoding/json"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	db  *sql.DB
	log logger.LoggerHandler
	m   *model
}

func newController(db *sql.DB, log logger.LoggerHandler) *controller {
	m := newModel(db, log)
	return &controller{db: db, log: log, m: m}
}

func (ctrl *controller) list(c *gin.Context) {
	examples, err := ctrl.m.list()
	if err != nil {
		c.JSON(err.(*vars.StatusErr).Code(), gin.H{
			"message": err.(*vars.StatusErr).Error(),
		})
		return
	}

	c.JSON(vars.STATUS_OK_CODE, examples)
}

func (ctrl *controller) create(c *gin.Context) {
	var example ExampleInput
	err := json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
			"message": vars.STATUS_SERVER_ERROR_STRING,
		})
		return
	}

	err = ctrl.m.create(c, example)
	if err != nil {
		c.JSON(err.(*vars.StatusErr).Code(), gin.H{
			"message": err.(*vars.StatusErr).Error(),
		})
		return
	}

	c.JSON(vars.STATUS_OK_CODE, gin.H{
		"message": vars.STATUS_OK_STRING,
	})
}

func (ctrl *controller) show(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
			"message": vars.STATUS_SERVER_ERROR_STRING,
		})
		return
	}

	example, err := ctrl.m.show(id)
	if err != nil {
		c.JSON(err.(*vars.StatusErr).Code(), gin.H{
			"message": err.(*vars.StatusErr).Error(),
		})
		return
	}

	c.JSON(vars.STATUS_OK_CODE, example)
}

func (ctrl *controller) update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
			"message": vars.STATUS_SERVER_ERROR_STRING,
		})
		return
	}

	var example ExampleInput
	err = json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
			"message": vars.STATUS_SERVER_ERROR_STRING,
		})
		return
	}

	err = ctrl.m.update(c, id, example)
	if err != nil {
		c.JSON(err.(*vars.StatusErr).Code(), gin.H{
			"message": err.(*vars.StatusErr).Error(),
		})
		return
	}

	c.JSON(vars.STATUS_OK_CODE, gin.H{
		"message": vars.STATUS_OK_STRING,
	})
}

func (ctrl *controller) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
			"message": vars.STATUS_SERVER_ERROR_STRING,
		})
		return
	}

	err = ctrl.m.delete(c, id)
	if err != nil {
		c.JSON(err.(*vars.StatusErr).Code(), gin.H{
			"message": err.(*vars.StatusErr).Error(),
		})
		return
	}

	c.JSON(vars.STATUS_OK_CODE, gin.H{
		"message": vars.STATUS_OK_STRING,
	})
}
