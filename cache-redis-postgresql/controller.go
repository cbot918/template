package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Controller struct {
	Db    *sql.DB
	Cache *redis.Client
}

func NewController(db *sql.DB, cache *redis.Client) *Controller {
	return &Controller{
		Db:    db,
		Cache: cache,
	}
}

func (ctr *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ctr *Controller) Health(c *gin.Context) {
	msg := ""
	if err := ctr.Db.Ping(); err != nil {
		msg += "db ping failed"
	}
	if res := ctr.Cache.Ping(context.Background()); res.Val() != "PONG" {
		msg += "cache ping failed"
	}
	if msg == "" {
		msg += "db cache ok"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func (ctr *Controller) HandleUserData(c *gin.Context) {

	log(c.Params.Get("id"))
}
