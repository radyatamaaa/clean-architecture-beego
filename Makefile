BINARY=engine
BRANCH=$(shell git rev-parse --abbrev-ref 'HEAD')
commit:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	git commit

test: 
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

unittest:
	go test -short  ./...

swagger_documentation:
	swag init -g main.go --output swagger

