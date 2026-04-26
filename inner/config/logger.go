package config

import (
	"log/slog"
	"os"
	"sync"

	"github.com/lmittmann/tint"
)

var GetLogger = sync.OnceValue(func() *slog.Logger {
	if os.Getenv("GIN_MODE") == "debug" {
		logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
            Level: slog.LevelDebug,
        }))

        slog.SetDefault(logger)
		return logger
	}
	
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)
    return logger
})
