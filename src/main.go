package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/AviSantoso/go-image-service/logger"
	gin "github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var UNKNOWN_ERROR = "An unknown error occurred."

func Hello(name string) (string, error) {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main/hello", id)
	message := fmt.Sprintf("Hello, %s!", name)
	log.Info(message)
	return message, nil
}

func HandlerHello(ctx *gin.Context) {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main/handler/hello", id)

	name := ctx.DefaultQuery("name", "there")
	message, err := Hello(name)

	if err != nil {
		log.Error(err.Error())
		CtxError(ctx)
		return
	}

	CtxOk(ctx, message)
}

func HandlerImageUpload(ctx *gin.Context) {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "handler/image/upload", id)
	file, err := ctx.FormFile("file")

	if err != nil {
		log.Error(err.Error())
		CtxError(ctx)
		return
	}

	content_type := (file.Header.Get("Content-Type"))

	if !strings.HasPrefix(content_type, "image/") {
		err := fmt.Sprintf("The content type %s is not supported for this route.", content_type)
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	log.Info(fmt.Sprintf("Received new upload request with filename %s and size %d", file.Filename, file.Size))
	out_file := filepath.Join(".", "temp", id)

	log.Info(fmt.Sprintf("Writing image to %s", out_file))

	err = ctx.SaveUploadedFile(file, out_file)

	if err != nil {
		log.Error(err.Error())
		CtxError(ctx)
		return
	}

	message := fmt.Sprintf("'%s' uploaded!", file.Filename)
	log.Info(message)

	CtxOk(ctx, message)
}

func CtxOk(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func CtxError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": UNKNOWN_ERROR,
	})
}

func main() {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main", id)

	temp_dir := filepath.Join(".", "temp")
	err := os.MkdirAll(temp_dir, os.ModePerm)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating temp directory. %v", err))
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			image := v1.Group("/image")
			{
				image.POST("/upload", HandlerImageUpload)
			}
		}
	}

	r.GET("/", HandlerHello)
	r.Run()
}
