package rest

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"digger/inner/docker"
	"digger/inner/rest/dto"
	"digger/inner/rest/model"

	"github.com/gin-gonic/gin"
	"github.com/moby/moby/client"
)

func GetActiveContainers(c *gin.Context) {
	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

	got, err := doc.ContainerList(c, client.ContainerListOptions{
		Filters: make(client.Filters).Add("status", "running"),
	})
	if err != nil {
		HandleError(c, err, "failed to get container list")
        return
	}

	arr := make([]model.ActiveContainer, len(got.Items))
	for i, v := range got.Items {
		arr[i] = model.ActiveContainer{
			Id:   v.ID,
			Name: v.Names[0],
		}
	}

	c.JSON(http.StatusOK, arr)
}

func GetContainerLogs(c *gin.Context) {
	var params dto.GetLogs
	if err := c.BindUri(&params); err != nil {
		HandleError(c, err, "wrong params")
        return
	}

	if err := c.BindQuery(&params); err != nil {
		HandleError(c, err, "wrong params")
        return
	}

    if params.Limit <= 0 {
        params.Limit = 1000
    }

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

    slog.Debug("params", "limit", params.Limit)
	res, err := doc.ContainerLogs(c, params.ID, client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
        Tail: strconv.Itoa(params.Limit),
	})
	if err != nil {
		HandleError(c, err, "failed to read logs")
        return
	}

    // TODO: parse logs format
    c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "ttachment; filename=\"logs.txt\"")
	_, err = io.Copy(c.Writer, res)
	if err != nil {
		slog.Error("failed to write logs", "err", err)
	}
}

func CollectRoutes(e *gin.RouterGroup) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	e.GET("/containers", GetActiveContainers)
	e.GET("/containers/:id/logs", GetContainerLogs)
}
