package main

import (
	"fmt"
	"net/http"

	logger "github.com/AviSantoso/go-image-service/logger"
	gin "github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func HelloHandler(c *gin.Context) {
	id, _ := gonanoid.New()
	log := logger.New("handler/hello", id)
	name := c.DefaultQuery("name", "there")

	message := fmt.Sprintf("Hello, %s!", name)
	log.Info(message)

	c.String(http.StatusOK, message)
}

func main() {
	r := gin.Default()
	r.GET("/hello", HelloHandler)
	r.Run()
}
