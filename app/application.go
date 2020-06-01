package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication will start the app
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
