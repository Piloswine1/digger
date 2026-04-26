package rest

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"digger/inner/config"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var f embed.FS

func getPrefix() string {
	uiPath := os.Getenv("UI_PATH")
	if uiPath != "" {
		return uiPath
	}

	return "/ui"
}

func UpUI(e *gin.Engine) {
	prefix := getPrefix()
	g := e.Group(prefix,
		config.Auth(),
		gzip.Gzip(gzip.BestCompression))

	e.LoadHTMLFS(http.FS(f), "dist/index.tmpl")
	g.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dist/index.tmpl", gin.H{
			"prefix": prefix,
		})
	})

	assets, err := fs.Sub(f, "dist")
	if err != nil {
		panic(err)
	}
	g.StaticFS("/", http.FS(assets))
}
