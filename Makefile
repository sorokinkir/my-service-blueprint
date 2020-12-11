#!/usr/bin/env bash

build:
	go build -o /bin/serviceAPI cmd/api/main.go

run:
	cd cmd/api; ./runserver.sh