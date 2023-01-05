package main

import (
	"github.com/gin-gonic/gin"
)

// En la funcion main se encuentran todas las rutas que tiene nuestra api
// Que tipo de autentificacion ocuparan y su handler.
func main() {
	router := gin.Default()

	router.GET("/", Healthcheck)
	router.POST("/test", Test)

	router.Run("0.0.0.0:8081")
}
