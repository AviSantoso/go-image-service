package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	imageservice "github.com/AviSantoso/go-image-service/image-service"
	"github.com/AviSantoso/go-image-service/logger"
	gin "github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var imageService = imageservice.New(os.Stdout)

func HandlerHello(ctx *gin.Context) {
	id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "main/handler/hello", id)

	name := ctx.DefaultQuery("name", "there")
	message, err := Hello(name)

	if err != nil {
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	CtxOk(ctx, message)
}

func HandlerImageUpload(ctx *gin.Context) {
	_id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "handler/image/upload", _id)

	id := ctx.Query("id")
	if len(id) == 0 {
		err := "Id must be a non-empty string."
		log.Error(err)
		CtxError(ctx, err)
		return
	}

	form_file, err := ctx.FormFile("file")
	if err != nil {
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	content_type := (form_file.Header.Get("Content-Type"))

	if !strings.HasPrefix(content_type, "image/") {
		err := fmt.Sprintf("The content type %s is not supported for this route.", content_type)
		log.Error(err)
		CtxError(ctx, err)
		return
	}

	file, err := form_file.Open()
	if err != nil {
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	res := imageService.UploadImage(id, bytes)

	ctx.JSON(res.Status, res.Data)
}

func getEngine() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			image := v1.Group("/image")
			{
				image.POST("/upload/:id", HandlerImageUpload)
			}
		}
	}

	router.GET("/", HandlerHello)

	return router
}
