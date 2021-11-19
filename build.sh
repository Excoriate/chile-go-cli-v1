#!/usr/bin/env bash

log () {
	echo "[$(date '+%H:%M:%S')] $1";
}

log "Building docker image for ChileCLI..."

docker build -t chilecli .

log "Docker image built."
log "Executing docker container..."

docker run -d -it chilecli

log "Successfully run"
docker ps -a
