
GOPATH:=$(shell go env GOPATH)

.PHONY: swag
swag:
	swag init --parseDependency --parseInternal --parseDepth 1

.PHONY: pull
pull:
	git pull

.PHONY: build
build:
	go env -w GOPROXY=https://goproxy.cn,direct
	go build -o run.sh

.PHONY: start
start:
	nohup ./run.sh >> logs/run.log 2>&1 & echo $! > run.PID

.PHONY: stop
stop:
	kill -9 cat run.PID

.PHONY: test
test:
	go test -v ./... -cover
