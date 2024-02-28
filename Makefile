.PHONY:build
build:
	go build -o jellyfin_exporter ./main.go

.PHONY:build_image
build_image:
    docker build -t jellyfin_exporter .

.DEFAULT_GOAL := build