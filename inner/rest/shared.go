package rest

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
)

func HandleError(c *gin.Context, err error, msg string) {
    slog.Error(msg, err)

    errF := eris.ToJSON(eris.Wrap(err, msg), true)
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": errF,
    })
}
