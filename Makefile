RELEASE_TYPE ?= patch 
# patch, minor, major
SEMVER_IMAGE ?= alpine/semver
# https://hub.docker.com/r/alpine/semver

all: test build release test-release clean
.PHONY: all

NEXT_VERSION := v$(shell docker run --rm $(SEMVER_IMAGE) semver -c -i $(RELEASE_TYPE) $(CURRENT_VERSION))
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
PWD = $(pwd)

test:
	go test -race -v ./...

build: main.go go.mod go.sum
	go build

release: check-version ~/.config/goreleaser/gitea_token
	if [ "$(BRANCH)" = "main" ];then \
	  git tag $(NEXT_VERSION) ;\
	  goreleaser check ;\
	  goreleaser release --rm-dist ;\
	  git push --tags ;\
	fi

test-release:
	goreleaser check
	goreleaser release --snapshot --rm-dist

test-environment:
	docker run -it --rm -w "$(PWD)" -v "$(PWD)":"$(PWD)" -dp 8080:8080 cosmtrek/air

clean:
	rm -rf dist tmp node_modules yarn.lock
	go clean -modcache
	go get -u && go mod tidy