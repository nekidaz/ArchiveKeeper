package router

import (
	"ZipArchive/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", controllers.MainPage)

	r.POST("/upload", controllers.UploadFiles)

	return r
}
