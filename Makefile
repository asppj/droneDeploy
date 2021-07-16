project=droneDeploy
mainFile=cmd/main.go

GO = CGO_ENABLED=0 GO111MODULE=on GOPROXY=https://goproxy.cn,direct go

BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
GO_VERSION := $(shell go version)
GIT_BRANCH := $(shell git symbolic-ref --short -q HEAD)
GIT_COMMIT_HASH := $(shell git rev-parse --verify HEAD)
GIT_COMMIT_MSG=$(shell git show -s --format=%ai) $(shell git show -s --format=%ae)
GO_FLAGS := -v -ldflags="-w -s -X 'github.com/asppj/droneDeploy/conf.Build=$(BUILD_DATE)' -X 'github.com/asppj/droneDeploy/conf.Commit=$(GIT_COMMIT_HASH)' -X 'github.com/asppj/droneDeploy/conf.Branch=$(GIT_BRANCH)' -X 'github.com/asppj/droneDeploy/conf.Message=${GIT_COMMIT_MSG}' -X 'github.com/asppj/droneDeploy/conf.Platform=${GO_VERSION}'"


build:
	$(GO) build $(GO_FLAGS) -o bin/runner  $(mainFile)


run:
	go run cmd.go