package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundMiddleware es un middleware que maneja rutas no encontradas
func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica si el código de estado es 404
		if c.Writer.Status() == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "La ruta que estás buscando no existe. Por favor verifica la URL.",
			})
		}
	}
}
