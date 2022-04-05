package example

import (
	"database/sql"
	"encoding/json"
	"go-api-starter-kit/config"
	"go-api-starter-kit/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	db  *sql.DB
	log *logger.Logger
	s   *service
}

func newController(db *sql.DB, log *logger.Logger) *controller {
	s := newService(db, log)
	return &controller{db: db, log: log, s: s}
}

func (ctrl *controller) list(c *gin.Context) {
	var err error
	var examples []Example = []Example{}

	aid, ok := c.MustGet("aid").(int)
	if ok {
		examples, err = ctrl.s.list(aid)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}
	} else {
		c.JSON(config.REQUEST_ERROR, gin.H{
			"message": config.REQUEST_STRING,
		})
		return
	}

	c.JSON(config.OK_STATUS, examples)
}

func (ctrl *controller) create(c *gin.Context) {
	var err error
	var example ExampleInput

	err = json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	aid, ok := c.MustGet("aid").(int)
	if ok {
		err = ctrl.s.create(c, example, aid)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}
	} else {
		c.JSON(config.REQUEST_ERROR, gin.H{
			"message": config.REQUEST_STRING,
		})
		return
	}

	c.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}

func (ctrl *controller) show(c *gin.Context) {
	var err error
	var example Example

	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	aid, ok := c.MustGet("aid").(int)
	if ok {
		example, err = ctrl.s.show(id, aid)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}
	} else {
		c.JSON(config.REQUEST_ERROR, gin.H{
			"message": config.REQUEST_STRING,
		})
		return
	}

	c.JSON(config.OK_STATUS, example)
}

func (ctrl *controller) update(c *gin.Context) {
	var err error

	id, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	var example ExampleInput
	err = json.NewDecoder(c.Request.Body).Decode(&example)
	if err != nil {
		ctrl.log.Error(err.Error())
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	aid, ok := c.MustGet("aid").(int)
	if ok {
		err = ctrl.s.update(c, id, example, aid)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}
	} else {
		c.JSON(config.REQUEST_ERROR, gin.H{
			"message": config.REQUEST_STRING,
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
		ctrl.log.Error(err.Error())
		c.JSON(config.SERVER_ERROR, gin.H{
			"message": config.SERVER_STRING,
		})
		return
	}

	aid, ok := c.MustGet("aid").(int)
	if ok {
		err = ctrl.s.delete(c, id, aid)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}
	} else {
		c.JSON(config.REQUEST_ERROR, gin.H{
			"message": config.REQUEST_STRING,
		})
		return
	}

	c.JSON(config.OK_STATUS, gin.H{
		"message": config.OK_STRING,
	})
}
