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
	router.PATCH("/test", Test)
	router.PATCH("/itemsV2/00040000000EACED00057708000110DB3DFF13B10000000EACED00057708000110D931861E0C0000000EACED00057708000110DB3DFF13B10000000EACED00057708000110D931861E0C", Test)

	router.Run("0.0.0.0:8080")
}
