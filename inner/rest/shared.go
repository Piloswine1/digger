package rest

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
)

func HandleError(c *gin.Context, err error, msg string) {
    slog.Error(msg, err)
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": eris.Wrap(err, msg).Error(),
    })
}
