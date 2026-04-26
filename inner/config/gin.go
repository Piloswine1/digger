package config

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var GetGin = sync.OnceValue(func() *gin.Engine {
	r := gin.New()
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	return r
})
