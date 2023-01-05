package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type boroughInterface struct {
	Borough string `json:"alcaldia"`
}

func Test(c *gin.Context) {
	// Body Http
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	fmt.Println(reqBody)
	c.JSON(http.StatusOK, reqBody)
}

func Healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "Ok 200")
}
