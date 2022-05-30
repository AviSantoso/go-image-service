package main

import (
	"io/ioutil"
	"os"

	"github.com/AviSantoso/go-image-service/logger"
	gin "github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func HandlerHello(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "there")
	message := Hello(name)
	CtxOk(ctx, message)
}

func HandlerImageUpload(ctx *gin.Context) {
	_id, _ := gonanoid.New()
	log := logger.New(os.Stdout, "handler/image/upload", _id)

	id := ctx.Param("id")
	form_file, err := ctx.FormFile("file")
	if err != nil {
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	file, err := form_file.Open()
	if err != nil {
		// notest
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		// notest
		log.Error(err.Error())
		CtxErrorUnknown(ctx)
		return
	}

	res := imageService.UploadImage(id, bytes)
	ctx.JSON(res.Status, res.Data)
}

func HandlerImageDownload(ctx *gin.Context) {
	id := ctx.Param("id")
	res := imageService.DownloadImage(id)
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
				image.GET("/download/:id", HandlerImageDownload)
			}
		}
	}

	router.GET("/", HandlerHello)

	return router
}
