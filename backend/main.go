package main

import (
	"backend/middlewares"
	"backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(middlewares.ErrorMiddleware())
	r.Use(middlewares.NotFoundMiddleware())
	routers.RegisterRouters(r)

	// Iniciar servidor
	r.Run(":8088")
}
