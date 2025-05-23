
all: help

.PHONY: help
help: Makefile
	@echo
	@echo "Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


## setup: install needed packages
.PHONY: setup
setup:
	@brew install watchexec
	@go install github.com/vektra/mockery/v2@v2.50.0


## start: build and run local project
.PHONY: start
start:
	watchexec -r -e go --wrap-process session -- "go run ./main.go"

## mock: generate the mocks
.PHONY: mock
mock:
	mockery --all


## test: run unit tests
.PHONY: test
test:
	ENV=.env.test go test -v ./...


## build: build the app into a binary
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./run-app ./main.go
