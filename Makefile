SHELL:= /bin/bash

protoc:
	protoc -I urlrepo/ urlrepo/urlrepo.proto --go_out=plugins=grpc:urlrepo

run:
	go install && urlrepo

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o urlrepo .

docker:
	docker-compose build

.PHONY: protoc run build docker
