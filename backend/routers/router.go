package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) {
	RegisterProductsRouter(router)
}
