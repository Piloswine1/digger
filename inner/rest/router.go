package rest

import (
	"log/slog"
	"net/http"
	"strconv"

	"digger/inner/config"
	"digger/inner/docker"
	"digger/inner/rest/dto"
	"digger/inner/rest/model"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/client"
)

func GetActiveContainers(c *gin.Context) {
    params := dto.ContainersReq{}
    if err := params.FromReq(c); err != nil {
        HandleError(c, err, "failed to get containers req")
        return
    }

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

	got, err := doc.ContainerList(c, client.ContainerListOptions{
        All: params.All,
    })
	if err != nil {
		HandleError(c, err, "failed to get container list")
		return
	}

	arr := make([]model.ContainerInfo, len(got.Items))
	for i, v := range got.Items {
		arr[i] = model.ContainerInfo{
			Id:     v.ID,
			Name:   v.Names[0],
			Status: v.Status,
		}
	}

	c.JSON(http.StatusOK, arr)
}

func GetContainerLogs(c *gin.Context) {
	params := dto.GetLogs{}
	if err := params.FromReq(c); err != nil {
		HandleError(c, err, "wrong params")
		return
	}

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

	slog.Debug("reading logs",
		"id", params.ID,
		"limit", params.Limit,
		"stderr", params.StdErr)
	res, err := doc.ContainerLogs(c, params.ID, client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: params.StdErr,
		Tail:       strconv.Itoa(params.Limit),
	})
	if err != nil {
		HandleError(c, err, "failed to read logs")
		return
	}

	// TODO: parse logs format
	c.Header("Content-Type", "text/plain")
	c.Header("Content-Disposition", "attachment; filename=\"logs.txt\"")
	_, err = stdcopy.StdCopy(c.Writer, c.Writer, res)
	if err != nil {
		slog.Error("failed to write logs", "err", err)
	}
}

func RestartContainer(c *gin.Context) {
    params := dto.ContainerID{}
    if err := params.FromReq(c); err != nil {
        HandleError(c, err, "failed to get containers req")
        return
    }

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

    slog.Debug("restarting container", "id", params.ID)
    _, err := doc.ContainerRestart(c, params.ID, client.ContainerRestartOptions{})
    if err != nil {
        HandleError(c, err, "failed to restart container")
        return
    }

	c.String(http.StatusOK, "ok")
}

func StopContainer(c *gin.Context) {
    params := dto.ContainerID{}
    if err := params.FromReq(c); err != nil {
        HandleError(c, err, "failed to get containers req")
        return
    }

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

    slog.Debug("stopping container", "id", params.ID)
    _, err := doc.ContainerStop(c, params.ID, client.ContainerStopOptions{})
    if err != nil {
        HandleError(c, err, "failed to stop container")
        return
    }

	c.String(http.StatusOK, "ok")
}

func StartContainer(c *gin.Context) {
    params := dto.ContainerID{}
    if err := params.FromReq(c); err != nil {
        HandleError(c, err, "failed to get containers req")
        return
    }

	doc := docker.MustNewDockerClient(c)
	defer doc.Close()

    slog.Debug("starting container", "id", params.ID)
    _, err := doc.ContainerStart(c, params.ID, client.ContainerStartOptions{})
    if err != nil {
        HandleError(c, err, "failed to start container")
        return
    }

	c.String(http.StatusOK, "ok")
}

func CollectRoutes(e *gin.RouterGroup) {
	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g := e.Group("/containers",
		config.Auth(),
		gzip.Gzip(gzip.DefaultCompression))
	g.GET("", GetActiveContainers)
	g.GET(":id/logs", GetContainerLogs)

	g.POST(":id/restart", RestartContainer)
	g.POST(":id/stop", StopContainer)
	g.POST(":id/start", StartContainer)
}
