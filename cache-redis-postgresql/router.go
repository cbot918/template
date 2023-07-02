package main

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, dep *Dep) *gin.Engine {
	ctr := NewController(dep.Db, dep.Cache)

	r.GET("/ping", ctr.Ping)
	r.GET("/health", ctr.Health)
	r.GET("/names/:id", ctr.HandleUserData)

	return r
}
