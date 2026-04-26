package docker

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moby/moby/client"
	"github.com/rotisserie/eris"
)

var NewDockerClient = func() (*client.Client, error) {
	return client.New(
		client.FromEnv,
		client.WithUserAgent("digger/0.0.1"),
	)
}

var MustNewDockerClient = func(c *gin.Context) *client.Client {
    slog.Debug("attemp to get docker",
        "DOCKER_HOST", os.Getenv("DOCKER_HOST"))
    client, err := client.New(
		client.FromEnv,
		client.WithUserAgent("digger/0.0.1"),
	)
    if err != nil {
        msg := "failed to get docker client"
        slog.Error(msg)
        c.JSON(
            http.StatusInternalServerError,
            eris.Wrap(err, msg))
    }
    return client
}
