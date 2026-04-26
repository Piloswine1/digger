#!/bin/env bash

podman run --rm \
	-p 8080:8080 \
	-v /run/user/1000/podman/podman.sock:/tmp/docker.sock:ro \
	-e GIN_MODE=release \
	-e DOCKER_HOST=unix:///tmp/docker.sock \
	digger
