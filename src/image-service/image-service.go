package imageservice

import (
	"fmt"
	"io"
	"net/http"

	"github.com/AviSantoso/go-image-service/logger"
	"github.com/AviSantoso/go-image-service/storage"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ImageService struct {
	storage storage.StorageInterface
	logger  logger.Logger
}

type HandlerResult struct {
	status int
	data   gin.H
}

func New(writer io.Writer) ImageService {
	id, _ := gonanoid.New()
	logger := logger.New(writer, "image-service/image-service", id)
	storage := storage.NewInMemoryStorage(writer)
	service := ImageService{storage: storage, logger: logger}
	return service
}

/** Uploads an image to the service using an id and byte array */
func (service ImageService) UploadImage(id string, bytes []byte) HandlerResult {
	service.logger.Info(fmt.Sprintf("Uploading an image with the id '%s'", id))

	err := service.storage.Store(id, bytes)
	if err != nil {
		service.logger.Error(err.Error())
		return HandlerResult{status: http.StatusBadRequest, data: gin.H{
			"message": err.Error(),
		}}
	}

	return HandlerResult{status: http.StatusOK, data: gin.H{
		"id":      id,
		"message": fmt.Sprintf("Successfully uploaded image with id %s", id),
	}}
}

/** Downloads an image from the service with the given id */
func (service ImageService) DownloadImage(id string) HandlerResult {
	service.logger.Info(fmt.Sprintf("Downloading an image with the id '%s'", id))
	image, err := service.storage.Retrieve(id)

	if err != nil {
		service.logger.Error(err.Error())
		return HandlerResult{status: http.StatusBadRequest, data: gin.H{
			"message": err.Error(),
		}}
	}

	return HandlerResult{status: http.StatusOK, data: gin.H{
		"id":      id,
		"message": fmt.Sprintf("Retrieved image with id %s", id),
		"image":   image,
	}}
}
