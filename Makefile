# application meta info
NAME := chio
VERSION= 0.0.1
REVISION := $(shell git rev-parse --short HEAD)
BUILDDATE := $(shell date '+%Y/%m/%d %H:%M:%S %Z')
GOVERSION := $(shell go version)
LDFLAGS := -X 'main.revision=$(REVISION)' -X 'main.version=$(VERSION)' -X 'main.buildDate=$(BUILDDATE)' -X 'main.goVersion=$(GOVERSION)'
ENTRYPOINT := cmd/main.go
REQUIRETESTPKG := pkg

all: dep production

dep:
	dep ensure

# production mode: make [production | pro | p]
production pro p: build-production test-production

# development mode: make [development | develop | dev | d]
development develop dev d: dep build-development

# buid
build-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-linux-64 ./$(ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-darwin-64 ./$(ENTRYPOINT)

# test
test-%:
	go test -tags="$* netgo" -count=1 ./$(REQUIRETESTPKG)

update:
	dep ensure -update

# publicing
deploy:
	mv ./bin/* ./public/
