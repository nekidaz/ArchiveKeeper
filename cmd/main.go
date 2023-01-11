package main

import (
	"ZipArchive/router"
)

func main() {
	r := router.SetupRouters()
	r.Run(":8080")
}
