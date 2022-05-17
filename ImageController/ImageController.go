package ImageController

import (
	"mime/multipart"

	gin "github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context, file *multipart.FileHeader) {
	// id, _ := gonanoid.New()
	// log := logger.New("handler/image/upload", id)

	// log.Info(fmt.Sprintf("Received new upload request with filename %s and size %d", file.Filename, file.Size))
	// out_file := filepath.Join(".", "temp", id)

	// http.DetectContentType(&file.content)

	// err = ctx.SaveUploadedFile(file, out_file)

	// if err != nil {
	// 	log.Error(err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "An unknown error occurred.",
	// 	})
	// 	return

	// }

	// message := fmt.Sprintf("'%s' uploaded!", file.Filename)
	// log.Info(message)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": message,
	// })
}
