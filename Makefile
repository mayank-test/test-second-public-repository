.PHONY: generate

# Let Go know that our modules are private
export GOPRIVATE=github.com/watchtowerai

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# SERVICE_NAME should contain the name of your service. The name 'service_template' is placed for test purposes only, so
# 'gateway' or whatever service name is expected here.
SERVICE_NAME=service_template
BINARY_NAME=./$(SERVICE_NAME)

# docker parameters
NAME=watchtowerai/$(SERVICE_NAME)
TAG=$(shell git log -1 --pretty=format:"%H")
VERSION=$(NAME):$(TAG)
LATEST=$(NAME):latest
GIT_USER?=""
GITHUB_OAUTH_TOKEN?=""
GO_ENV?=test
GO_GIN_MODE?=release

all: clean build start
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	GIN_MODE=$(GO_GIN_MODE) GO_ENV=$(GO_ENV) $(GOTEST) ./... -count=1 -coverprofile=coverage.out
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
dockerbuild:
	docker build -t $(VERSION) . --build-arg GIT_USER=$(GIT_USER) --build-arg GITHUB_OAUTH_TOKEN=$(GITHUB_OAUTH_TOKEN) --build-arg SERVICE_NAME=$(SERVICE_NAME)
	docker tag $(VERSION) $(LATEST)
dockertag:
	docker tag $(SERVICE_NAME):latest $(VERSION)
	docker tag $(SERVICE_NAME):latest $(LATEST)
dockerpush:
	docker push $(VERSION)
	docker push $(LATEST)
start:
	./$(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	go mod download
generate:
	go generate ./...

