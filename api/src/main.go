package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	imageservice "github.com/AviSantoso/go-image-service/image-service"
	"github.com/AviSantoso/go-image-service/logger"
	"github.com/AviSantoso/go-image-service/storage"
	gin "github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var UNKNOWN_ERROR = "An unknown error occurred."
var imageService = imageservice.New(os.Stdout, storage.NewInMemoryStorage(os.Stdout))

func Hello(name string) string {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main/hello", id)
	message := fmt.Sprintf("Hello, %s!", name)
	log.Info(message)
	return message
}

func CtxOk(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func CtxError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

func CtxErrorUnknown(ctx *gin.Context) {
	CtxError(ctx, UNKNOWN_ERROR)
}

func main() {
	// notest
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main", id)

	temp_dir := filepath.Join(".", "temp")
	err := os.MkdirAll(temp_dir, os.ModePerm)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating temp directory. %v", err))
	}

	engine := getEngine()
	engine.Run()
}
