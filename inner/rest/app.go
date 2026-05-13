package rest

import (
	"digger/inner/config"
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func InitApp() {
	r := config.GetGin()

    logger := config.GetLogger()
	r.Use(sloggin.New(logger))
	r.Use(gin.Recovery())

    CollectUI(r)

    group := r.Group("/api/v1")
	CollectRoutes(group)

    slog.Info("starting server as 0.0.0.0:8080")
	if err := r.Run(); err != nil {
		slog.Debug("failed to run", "err", err)
		panic(err)
	}
}
