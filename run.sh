#!/usr/bin/env bash

log () {
	echo "[$(date '+%H:%M:%S')] $1";
}

log "Compiling CLI..."
GOGC=off go build -o chilecli main.go

log "Executing chile-cli"
./chilecli "$@"
