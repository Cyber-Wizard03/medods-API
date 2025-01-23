package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const MAX_UPLOAD_SIZE = 20

var IMAGE_TYPES = map[string]interface{}{
	"image/jpeg": nil,
	"image/png":  nil,
}

func Upload(c *gin.Context) {

	form, _ := c.MultipartForm()
	fmt.Println(form)

	image := form.File["photo"]
	pdf := form.File["pdf"]
	word := form.File["word"]
	cover := form.File["cover"]

	for _, file := range image {
		log.Println("image:", file.Filename)
		c.SaveUploadedFile(file, "./../frontend/public/pictures/"+file.Filename)
	}
	for _, file := range pdf {
		log.Println("pdf:", file.Filename)
		c.SaveUploadedFile(file, "./../frontend/public/pdf/"+file.Filename)
	}
	for _, file := range word {
		log.Println("word:", file.Filename)
		c.SaveUploadedFile(file, "./../frontend/public/word/"+file.Filename)
	}
	for _, file := range cover {
		log.Println("cover:", file.Filename)
		c.SaveUploadedFile(file, "./../frontend/public/pictures/"+file.Filename)
	}
	fmt.Println("4")
	c.String(http.StatusOK, "Данные добавлены")
}
