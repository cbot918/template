package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// read config
	cfg := NewConfig()

	// DI
	dep := NewDep()
	dep.Db = NewDb(cfg)
	dep.Cache = NewCache(cfg)

	// setup apiServer
	app := gin.Default()
	app = NewRouter(app, dep)

	// apiServer run
	app.Run(cfg.PORT)
}
