package rest

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func InitApp() {
	r := gin.New()

    slog.SetLogLoggerLevel(slog.LevelDebug)
	logger := slog.Default()
	r.Use(sloggin.New(logger))

	r.Use(gin.Recovery())

    group := r.Group("/v1")
	CollectRoutes(group)

    slog.Info("starting server as 0.0.0.0:8080")
	if err := r.Run(); err != nil {
		slog.Debug("failed to run", "err", err)
		panic(err)
	}
}
