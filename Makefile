DIR := $(shell mktemp -d)

install_deps:
	@echo "Installing go-swagger..."
	@go get -u github.com/swaggo/swag/cmd/swag
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Done"

swag:
	@~/go/bin/swag init --parseDependency --parseInternal

build:
	@go build -o bin/showandtell main.go

serve:
	@go run main.go --serve

static:
	@cd $(STATIC_DIR) && npm run build

compile:
	@echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -o build/showandtell-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o build/showandtell-linux-amd64 main.go
	GOOS=freebsd GOARCH=386 go build -o build/showandtell-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o build/showandtell-windows-386.exe main.go

all: build
