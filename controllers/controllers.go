package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
	"log"
	"net/http"
	"os"
	"strings"
)

func MainPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func UploadFiles(c *gin.Context) {
	file, _ := c.FormFile("file")
	maxSize := int64(100 * 1024 * 1024)

	if file.Size > maxSize {
		c.String(200, "Файл не должен быть больше 100мб")
		return
	}

	err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)

	if err != nil {
		log.Println("Error saving file: ", err)
		c.String(http.StatusInternalServerError, "Error saving file")
		return
	}
	
	parts := strings.Split(file.Filename, ".")

	if err := archiver.Archive([]string{"uploads/" + file.Filename}, "zipFiles/"+parts[0]+".zip"); err != nil {
		log.Fatalf("Ошибка при создании архива: %v", err)
	}

	c.File("zipFiles/" + parts[0] + ".zip")

	os.Remove("uploads/" + file.Filename)
	os.Remove("zipFiles/" + parts[0] + ".zip")
}
