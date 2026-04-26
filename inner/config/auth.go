package config

import (
	"log/slog"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

const DEFAULT_LOGIN = "diggerlogin"
const DEFAULT_PASS = "diggerpass"

var Auth = sync.OnceValue(func() gin.HandlerFunc {
    login := os.Getenv("USER_LOGIN")
    pass := os.Getenv("USER_PASS")

    if login == "" {
        slog.Warn("using insecure login", "val", DEFAULT_LOGIN)
        login = DEFAULT_LOGIN
    }
    if pass == "" {
        slog.Warn("using insecure pass", "val", DEFAULT_PASS)
        pass = DEFAULT_PASS
    }

    return gin.BasicAuth(gin.Accounts{
        login: pass,
    })
})
