package example

import (
	"database/sql"
	"encoding/json"
	"go-api-starter-kit/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type controller struct {
	db  *sql.DB
	log *zap.Logger
	s   *service
}

func newController(db *sql.DB, log *zap.Logger) *controller {
	s := newService(db, log)
	return &controller{db: db, log: log, s: s}
}

func (ctrl *controller) list(c *gin.Context) {
	examples, err := ctrl.s.list()
	if err != nil {
		c.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	c.JSON(config.OK_STATUS, examples)
}

func (ctrl *controller) create(c *gin.Context) {
	var example examplereq
	err := json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(config.SERVER_STRING, zap.Error(err))
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	err = ctrl.s.create(c, example)
	if err != nil {
		c.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	c.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}

func (ctrl *controller) show(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(config.SERVER_STRING, zap.Error(err))
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	example, err := ctrl.s.show(id)
	if err != nil {
		c.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	c.JSON(config.OK_STATUS, example)
}

func (ctrl *controller) update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(config.SERVER_STRING, zap.Error(err))
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	var example examplereq
	err = json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(config.SERVER_STRING, zap.Error(err))
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	err = ctrl.s.update(c, id, example)
	if err != nil {
		c.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	c.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}

func (ctrl *controller) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(config.SERVER_STRING, zap.Error(err))
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	err = ctrl.s.delete(c, id)
	if err != nil {
		c.JSON(err.(*config.CustErr).Code(), gin.H{
			"message": err.(*config.CustErr).Error(),
		})
		return
	}
	c.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}
